package spotify

import (
	"errors"
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
	urlParts := strings.Split(url.Path, "/")
	// looking for a URL that looks something like
	// /v1/playlists
	if len(urlParts) < 2 {
		return errors.New("invalid endpoint structure")
	}
	apiVersion := urlParts[0]
	path := strings.Join(urlParts[1:], "/")
	return api.get(apiVersion, path, url.Query(), obj)
}

// URL parses HREF into a *net/url.URL
func (h *HREF) URL() (*url.URL, error) {
	return url.Parse(string(*h))
}
