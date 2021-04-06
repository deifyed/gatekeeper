package discovery

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func FetchDiscoveryDocument(discoveryURL url.URL) (Document, error) {
	response, err := http.Get(discoveryURL.String())
	if err != nil {
		return Document{}, fmt.Errorf("fetching discovery URL: %w", err)
	}

	doc := Document{}

	err = json.NewDecoder(response.Body).Decode(&doc)
	if err != nil {
		return Document{}, fmt.Errorf("decoding json response: %w", err)
	}

	return doc, nil
}
