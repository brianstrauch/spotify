package spotify

import (
	"net/url"
	"strings"
)

type HREF string

// Get returns the most recent version of the object in its entirety
func (h *HREF) Get(api *API, obj interface{}) error {
	url, err := h.URL()
	if err != nil {
		return err
	}

	// Example path: v1/me/player
	idx := strings.Index(url.Path, "/")
	version := url.Path[:idx]
	endpoint := url.Path[idx:]

	return api.get(version, endpoint, url.Query(), obj)
}

// URL parses HREF into a URL object
func (h *HREF) URL() (*url.URL, error) {
	return url.Parse(string(*h))
}
