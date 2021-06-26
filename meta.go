package spotify

import (
	"net/url"
	"strings"
)

type HREF string

// Get returns the most recent version of the object in it's entirety
func (h *HREF) Get(api *API, obj interface{}) error {
	url, err := h.GetURL()
	if err != nil {
		return err
	}
	path := strings.Replace(url.Path, "/v1", "", 1)
	return api.get(path, obj)
}

// GetURL parses HREF into a *net/url.URL
func (h *HREF) GetURL() (*url.URL, error) {
	return url.Parse(string(*h))
}

// Meta represents common fields found in most API responses
type Meta struct {
	// HREF is the URL for fetching the entire resource from the API
	HREF *HREF
	// ExternalURLs contains external URLs relating to the object
	// It's only key is `spotify` as of 2021-06-26
	ExternalURLs map[string]string `json:"external_urls"`
	ID           string            `json:"id"`
	// Type represents the type of the object.
	// Common types are 'album', 'playlist', and 'track'
	Type string `json:"type"`
	// URI is the Spotify URI for the object
	URI string `json:"uri"`
}

// Get returns the most recent version of the object in it's entirety
func (m *Meta) Get(api *API, obj interface{}) error {
	return m.HREF.Get(api, obj)
}

// ItemsMeta represents common fields in responses with an `items` field
type ItemsMeta struct {
	// HREF is the URL for fetching the entire resource from the API
	HREF *HREF
	// Limit is the limit provided in the original request
	Limit int `json:"limit"`
	// Next is the API URl for the next page of the response
	Next *string `json:"next"`
	// Offset is the offset in the original request
	Offset int `json:"offset"`
	// Previous is the API URL of the previous page of the response
	Previous *string `json:"previous"`
	// Total is the total number of objects in the response
	Total int `json:"total"`
}

// Get returns the most recent version of the object in it's entirety
func (im *ItemsMeta) Get(api *API, obj interface{}) error {
	return im.HREF.Get(api, obj)
}
