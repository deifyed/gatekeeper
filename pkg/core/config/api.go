package config

import (
	"net/url"
	"os"
	"strings"
)

func LoadConfig() Config {
	get := generateGetter(os.Getenv)
	lGet := generateListGetter(os.Getenv, ";")

	cfg := Config{
		Port:         get("PORT", "4554"),
		ClientID:     get("CLIENT_ID", ""),
		ClientSecret: get("CLIENT_SECRET", ""),
		CookiePrefix: get("COOKIE_PREFIX", ""),
		Origins:      lGet("ALLOWED_ORIGINS", []string{}),
	}

	if baseURL, _ := url.Parse(get("BASE_URL", "")); baseURL != nil {
		cfg.BaseURL = *baseURL
	}

	if discoveryURL, _ := url.Parse(get("DISCOVERY_URL", "")); discoveryURL != nil {
		cfg.DiscoveryURL = *discoveryURL
	}

	return cfg
}

func generateGetter(getter stringGetter) func(string, string) string {
	return func(key string, defaultValue string) string {
		potentialValue := getter(key)

		if potentialValue == "" {
			return defaultValue
		}

		return potentialValue
	}
}

func generateListGetter(getter stringGetter, delimiter string) func(string, []string) []string {
	return func(key string, defaultValue []string) []string {
		potentialValue := getter(key)

		if potentialValue == "" {
			return defaultValue
		}

		cleanedValue := strings.Trim(potentialValue, delimiter)
		cleanedValue = strings.ReplaceAll(cleanedValue, "\n", "")
		cleanedValue = strings.ReplaceAll(cleanedValue, " ", "")

		parts := strings.Split(cleanedValue, delimiter)

		return parts
	}
}
