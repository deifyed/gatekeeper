package oauth2

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type openIDConfiguration struct {
	AuthorizationEndpoint string `json:"authorization_endpoint"`
	TokenEndpoint         string `json:"token_endpoint"`
	UserinfoEndpoint      string `json:"userinfo_endpoint"`
	EndSessionEndpoint    string `json:"end_session_endpoint"`
}

func Discover(discoveryURL string) (openIDConfiguration, error) {
	request, err := http.NewRequest(http.MethodGet, discoveryURL, nil)
	if err != nil {
		return openIDConfiguration{}, fmt.Errorf("preparing request: %w", err)
	}

	client := http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return openIDConfiguration{}, fmt.Errorf("making request: %w", err)
	}

	var cfg openIDConfiguration

	err = json.NewDecoder(response.Body).Decode(&cfg)
	if err != nil {
		return openIDConfiguration{}, fmt.Errorf("decoding response: %w", err)
	}

	return cfg, nil
}
