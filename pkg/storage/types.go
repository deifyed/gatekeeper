package storage

import (
	"errors"
	"net/url"
)

type StateClient interface {
	DeleteState(id string) (err error)
	GetState(id string) (state string, err error)
	PutState(id, state string) (err error)
}

type RedirectClient interface {
	DeleteRedirect(id string) (err error)
	GetRedirect(id string) (redirectURL url.URL, err error)
	PutRedirect(id string, redirectURL url.URL) (err error)
}

type Client interface {
	StateClient
	RedirectClient
}

// ErrNotFound indicates a missing resource
var ErrNotFound = errors.New("not found")
