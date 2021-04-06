package memory

import (
	"fmt"
	"github.com/deifyed/gatekeeper/pkg/state"
)

func New() state.Storage {
	return &stateClient{state: map[string]string{}}
}

func (s stateClient) Delete(id string) (err error) {
	delete(s.state, id)

	return nil
}

func (s stateClient) Get(id string) (string, error) {
	stateValue, ok := s.state[id]
	if !ok {
		return "", fmt.Errorf("getting state with id %s: %w", id, state.StorageErrNotFound)
	}

	return stateValue, nil
}

func (s stateClient) Put(id, state string) (err error) {
	s.state[id] = state

	return nil
}
