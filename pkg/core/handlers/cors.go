package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func allowOrigin(c *gin.Context, origin string) {
	c.Header("Access-Control-Allow-Origin", origin)
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	c.Header("Access-Control-Allow-Credentials", "true")

	c.Status(http.StatusNoContent)
}

func CreateCORSHandler(allowedOrigins []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		currentOrigin := c.Request.Header.Get("Origin")

		for _, item := range allowedOrigins {
			if item == currentOrigin {
				allowOrigin(c, currentOrigin)

				return
			}
		}

		c.Status(http.StatusBadRequest)
	}
}
