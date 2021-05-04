package memory

import (
	"fmt"
	"net/url"

	"github.com/deifyed/gatekeeper/pkg/storage"
)

func New() storage.Client {
	return &memoryClient{
		state:    map[string]string{},
		redirect: map[string]url.URL{},
	}
}

func (m *memoryClient) DeleteState(id string) (err error) {
	delete(m.state, id)

	return nil
}

func (m *memoryClient) GetState(id string) (string, error) {
	stateValue, ok := m.state[id]
	if !ok {
		return "", fmt.Errorf("getting state with id %s: %w", id, storage.ErrNotFound)
	}

	return stateValue, nil
}

func (m *memoryClient) PutState(id, state string) (err error) {
	m.state[id] = state

	return nil
}

func (m *memoryClient) DeleteRedirect(id string) (err error) {
	delete(m.redirect, id)

	return nil
}

func (m *memoryClient) GetRedirect(id string) (redirectURL url.URL, err error) {
	redirectValue, ok := m.redirect[id]
	if !ok {
		return url.URL{}, storage.ErrNotFound
	}

	return redirectValue, nil
}

func (m *memoryClient) PutRedirect(id string, redirectURL url.URL) (err error) {
	m.redirect[id] = redirectURL

	return nil
}
