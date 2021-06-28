package spotify

import (
	"net/url"
	"strings"
)

type HREF string

// Get returns the most recent version of the object in it's entirety
func (h *HREF) Get(api *API, obj interface{}) error {
	url, err := h.URL()
	if err != nil {
		return err
	}
	path := strings.Replace(url.Path, "/v1", "", 1)
	return api.get(path, obj)
}

// URL parses HREF into a *net/url.URL
func (h *HREF) URL() (*url.URL, error) {
	return url.Parse(string(*h))
}
