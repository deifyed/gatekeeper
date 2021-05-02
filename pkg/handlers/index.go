package handlers

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

type status struct {
	SpecificationURL string `json:"specification_url"`
}

func CreateIndexHandler(baseURL url.URL) gin.HandlerFunc {
	specificationURL := baseURL
	specificationURL.Path = "/specification"

	statusResponse := status{
		SpecificationURL: specificationURL.String(),
	}

	return func(c *gin.Context) {
		c.JSON(http.StatusOK, statusResponse)
	}
}
