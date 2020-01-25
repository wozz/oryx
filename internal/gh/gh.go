package gh

import (
	"path"

	"github.com/google/go-github/github"
	"github.com/gregjones/httpcache"
	"github.com/gregjones/httpcache/diskcache"
	homedir "github.com/mitchellh/go-homedir"
)

var client *github.Client

// Construct client lazily such that we don't complain about
// credentials unless needed.
func Client() *github.Client {
	if client == nil {
		client = newGithubClient()
	}
	return client
}

// Create new client.
func newGithubClient() *github.Client {
	// Create a cache/transport implementation
	var cacheTransport *httpcache.Transport
	homeDir, err := homedir.Dir()
	if err == nil {
		cacheDir := path.Join(homeDir, ".cache", "oryx", "gh")
		cache := diskcache.New(cacheDir)
		cacheTransport = httpcache.NewTransport(cache)
	} else {
		// Couldn't get homedir, just use mem cache instead
		// though of limited utility...
		cacheTransport = httpcache.NewMemoryCacheTransport()
	}

	return github.NewClient(cacheTransport.Client())
}
