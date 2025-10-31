package api

import (
	"time"

	gh "github.com/cli/go-gh/v2/pkg/api"
)

func newCacheRESTClient() (*gh.RESTClient, error) {
	opts := gh.ClientOptions{
		CacheTTL:    time.Hour * 24,
		EnableCache: true,
		Timeout:     time.Second * 30,
	}

	return gh.NewRESTClient(opts)
}

func newRESTClient() (*gh.RESTClient, error) {
	opts := gh.ClientOptions{
		Timeout: time.Second * 30,
	}

	return gh.NewRESTClient(opts)
}
