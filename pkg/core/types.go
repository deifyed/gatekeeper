package core

import "errors"

type StateStorage interface {
	Delete(id string) (err error)
	Get(id string) (state string, err error)
	Put(id, state string) (err error)
}

var StateStorageErrNotFound = errors.New("not found")
