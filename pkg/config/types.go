package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/url"

	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Config struct {
	LogLevel logrus.Level
	BaseURL  url.URL
	Port     string

	DiscoveryURL url.URL
	ClientID     string
	ClientSecret string

	Origins   []string
	Upstreams map[string]string

	CookiePrefix string
}

func (c Config) Validate() error {
	if err := validation.Validate(c.BaseURL.String(), validation.Required, is.URL); err != nil {
		return fmt.Errorf("BASE_URL: %w", err)
	}

	if err := validation.Validate(c.DiscoveryURL.String(), validation.Required, is.URL); err != nil {
		return fmt.Errorf("DISCOVERY_URL: %w", err)
	}

	return validation.ValidateStruct(&c,
		validation.Field(&c.Port, validation.Required, is.Port),
		validation.Field(&c.Origins, validation.Each(is.URL)),
		validation.Field(&c.ClientID, validation.Required),
		validation.Field(&c.ClientSecret, validation.Required),
		validation.Field(&c.Upstreams, validation.Each(is.URL)),
	)
}

type stringGetter func(key string) (value string)
