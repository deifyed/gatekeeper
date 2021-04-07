package config

import (
	"github.com/sirupsen/logrus"
	"net/url"
	"os"
	"strings"
)

func LoadConfig() Config {
	get := generateGetter(os.Getenv)
	lGet := generateListGetter(os.Getenv, ";")

	cfg := Config{
		LogLevel:     stringToLogLevel(get("LOG_LEVEL", "info")),
		Port:         get("PORT", "4554"),
		ClientID:     get("CLIENT_ID", ""),
		ClientSecret: get("CLIENT_SECRET", ""),
		CookiePrefix: get("COOKIE_PREFIX", ""),
		Origins:      lGet("ALLOWED_ORIGINS", []string{}),
		Upstreams:    map[string]string{},
	}

	if baseURL, _ := url.Parse(get("BASE_URL", "")); baseURL != nil {
		cfg.BaseURL = *baseURL
	}

	if discoveryURL, _ := url.Parse(get("DISCOVERY_URL", "")); discoveryURL != nil {
		cfg.DiscoveryURL = *discoveryURL
	}

	rawUpstreams := lGet("UPSTREAMS", []string{})
	for _, upstream := range rawUpstreams {
		parts := strings.Split(upstream, "=")

		cfg.Upstreams[parts[0]] = parts[1]
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

func stringToLogLevel(rawLevel string) logrus.Level {
	switch rawLevel {
	case "debug":
		return logrus.DebugLevel
	case "warning":
		return logrus.WarnLevel
	case "error":
		return logrus.ErrorLevel
	default:
		return logrus.InfoLevel
	}
}
