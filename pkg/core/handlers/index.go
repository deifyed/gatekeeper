package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
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
