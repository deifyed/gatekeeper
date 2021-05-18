// Package handlers contains handler functions for all the Gatekeeper entrypoints
package handlers

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/deifyed/gatekeeper/pkg/storage"
	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

const authcodeflowPath = "pkg/handlers/authcodeflow.go"

// CreateLoginHandler -
func CreateLoginHandler(storage storage.Client, opts CreateLoginHandlerOpts) gin.HandlerFunc {
	logger := opts.Logger.WithFields(map[string]interface{}{
		"file": authcodeflowPath,
		"func": "CreateLoginHandler",
	})

	return func(c *gin.Context) {
		redirectURI, err := url.Parse(c.Query("redirect"))
		if err != nil {
			c.Status(http.StatusBadRequest)

			return
		}

		stateID := uuid.New().String()
		state := uuid.New().String()

		err = storage.PutRedirect(stateID, *redirectURI)
		if err != nil {
			c.Status(http.StatusInternalServerError)

			return
		}

		err = storage.PutState(stateID, state)
		if err != nil {
			logger.Error(fmt.Errorf("error storing state ID: %w", err))

			c.Status(http.StatusInternalServerError)

			return
		}

		opts.CookieHandler.SetStateID(c.Writer, stateID)

		c.Redirect(http.StatusFound, opts.Oauth2Config.AuthCodeURL(state))
	}
}

func CreateCallbackHandler(storage storage.Client, opts CreateCallbackHandlerOpts) gin.HandlerFunc {
	logger := opts.Logger.WithFields(map[string]interface{}{
		"file": authcodeflowPath,
		"func": "CreateCallbackHandler",
	})

	return func(c *gin.Context) {
		providedState := c.Query("state")
		code := c.Query("code")
		stateID, _ := opts.CookieHandler.GetStateID(c.Request)

		existingState, _ := storage.GetState(stateID)
		if providedState != existingState {
			logger.Error("invalid state provided")

			c.Status(http.StatusBadRequest)

			return
		}

		err := storage.DeleteState(stateID)
		if err != nil {
			logger.Error("error deleting state cookie", err)

			c.Status(http.StatusInternalServerError)

			return
		}

		token, err := opts.Oauth2Config.Exchange(opts.Ctx, code)
		if err != nil {
			logger.Error(err)
		}

		rawIDToken, ok := token.Extra("id_token").(string)
		if !ok {
			logger.Errorf("error extracting ID token")
		}

		idToken, err := opts.TokenVerifier.Verify(opts.Ctx, rawIDToken)
		if err != nil {
			logger.Error(err)
		}

		logger.Debug("found idToken: ", idToken)

		opts.CookieHandler.SetAccessToken(c.Writer, token.AccessToken, int(token.Expiry.Unix()))
		opts.CookieHandler.SetRefreshToken(c.Writer, token.RefreshToken)
		opts.CookieHandler.SetIDToken(c.Writer, rawIDToken)

		redirectURL, err := storage.GetRedirect(stateID)
		if err != nil {
			c.Status(http.StatusInternalServerError)

			return
		}

		c.Redirect(http.StatusFound, redirectURL.String())
	}
}

func CreateUserinfoHandler(opts CreateUserinfoHandlerOpts) gin.HandlerFunc {
	logger := opts.Logger.WithFields(map[string]interface{}{
		"file": authcodeflowPath,
		"func": "CreateUserinfoHandler",
	})

	return func(c *gin.Context) {
		tokenSource := newTokenGetter(opts.CookieHandler, c)

		userinfo, err := opts.Provider.UserInfo(opts.Ctx, tokenSource)
		if err != nil {
			logger.Warn(fmt.Errorf("error getting userinfo: %w", err))

			c.Status(http.StatusInternalServerError)

			return
		}

		c.JSON(http.StatusOK, userinfo)
	}
}
