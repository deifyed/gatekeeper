package memory

import "net/url"

type memoryClient struct {
	state    map[string]string
	redirect map[string]url.URL
}
