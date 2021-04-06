package memory

import (
	"fmt"

	"github.com/deifyed/gatekeeper/pkg/core"
)

func New() core.StateStorage {
	return &stateClient{state: map[string]string{}}
}

func (s stateClient) Delete(id string) (err error) {
	delete(s.state, id)

	return nil
}

func (s stateClient) Get(id string) (state string, err error) {
	state, ok := s.state[id]
	if !ok {
		return "", fmt.Errorf("getting state with id %s: %w", id, core.StateStorageErrNotFound)
	}

	return state, nil
}

func (s stateClient) Put(id, state string) (err error) {
	s.state[id] = state

	return nil
}
