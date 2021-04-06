package state

import "errors"

type Storage interface {
	Delete(id string) (err error)
	Get(id string) (state string, err error)
	Put(id, state string) (err error)
}

var StorageErrNotFound = errors.New("not found")
