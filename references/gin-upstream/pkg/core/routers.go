/*
 * Reference Upstream
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package core

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/coreos/go-oidc"
	"github.com/deifyed/gatekeeper/pkg/discovery"
	"github.com/gin-gonic/gin"
)

// NewRouter returns a new router.
func NewRouter() *gin.Engine {
	router := gin.Default()

	discoveryURL, err := url.Parse(os.Getenv("DISCOVERY_URL"))
	if err != nil {
		log.Fatal(fmt.Errorf("error parsing DISCOVERY_URL: %w", err))
	}

	document, err := discovery.FetchDiscoveryDocument(*discoveryURL)
	if err != nil {
		log.Fatal(fmt.Errorf("error fetching discovery URL: %w", err))
	}

	ctx := context.Background()
	clientID := os.Getenv("CLIENT_ID")

	// Unprotected routes before the middleware
	router.GET("/open", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{})
	})

	authMiddleware, err := createAuthMiddleware(ctx, document.Issuer, clientID)
	if err != nil {
		log.Fatal(fmt.Errorf("error creating auth middleware"))
	}

	router.Use(authMiddleware)

	// Protected routes after the middleware
	router.GET("/closed", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{})
	})

	return router
}

func createAuthMiddleware(ctx context.Context, issuer, clientID string) (gin.HandlerFunc, error) {
	provider, err := oidc.NewProvider(ctx, issuer)
	if err != nil {
		return nil, fmt.Errorf("creating provider: %w", err)
	}

	verifier := provider.Verifier(&oidc.Config{
		ClientID:             clientID,
		SupportedSigningAlgs: []string{oidc.RS256},
	})

	return func(c *gin.Context) {
		rawIDToken := c.GetHeader("x-id-token")
		if rawIDToken == "" {
			log.Println("no ID token found")

			c.AbortWithStatus(http.StatusUnauthorized)

			return
		}

		_, err := verifier.Verify(ctx, rawIDToken)
		if err != nil {
			log.Println(fmt.Errorf("failed to verify ID token: %w", err))

			c.AbortWithStatus(http.StatusUnauthorized)

			return
		}

		c.Next()
	}, nil
}
