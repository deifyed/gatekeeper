package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

const proxyHandlersFile = "pkg/handlers/proxy.go"

func CreateProxyHandler(opts CreateProxyHandlerOpts) gin.HandlerFunc {
	logger := opts.Logger.WithFields(map[string]interface{}{
		"file": proxyHandlersFile,
		"func": "CreateProxyHandler",
	})
	
	proxyMap := map[string]*httputil.ReverseProxy{}

	for name, rawTargetURL := range opts.Upstreams {
		upstreamURL, _ := url.Parse(rawTargetURL)

		proxyMap[name] = httputil.NewSingleHostReverseProxy(upstreamURL)
	}

	return func(c *gin.Context) {
		name := strings.Split(c.Request.URL.Path, "/")[1]
		proxy, ok := proxyMap[name]
		
		if !ok {
			c.Status(http.StatusNotFound)

			return
		}

		// TODO: Refresh token if necessary
		if accessToken, err := opts.CookieHandler.GetAccessToken(c); err != nil {
			c.Request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
		} else {
			logger.Warn(fmt.Errorf("fetching access token: %w", err))
		}
		
		if idToken, err := opts.CookieHandler.GetIDToken(c); err != nil {
			c.Request.Header.Add("x-id-token", idToken)
		} else {
			logger.Warn(fmt.Errorf("fetching ID token: %w", err))
		}
		
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
