package spotify

import (
	"net/url"
	"strconv"
	"strings"
	"time"
)

// Duration is a wrapper struct for time.Duration to unmarshal durations in milliseconds.
type Duration struct {
	time.Duration
}

func (d *Duration) UnmarshalJSON(data []byte) error {
	ms, err := strconv.Atoi(string(data))
	if err != nil {
		return err
	}
	d.Duration = time.Duration(ms * 1000000)
	return nil
}

type HREF string

// Get returns the most recent version of the object in its entirety.
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

// URL parses HREF into a URL object.
func (h *HREF) URL() (*url.URL, error) {
	return url.Parse(string(*h))
}

// Meta represents common fields found in most API responses.
type Meta struct {
	HREF         HREF
	ExternalURLs map[string]string `json:"external_urls"`
	ID           string            `json:"id"`
	Type         string            `json:"type"`
	URI          string            `json:"uri"`
}

// PagingMeta represents common fields in paged responses
type PagingMeta struct {
	HREF     HREF
	Limit    int    `json:"limit"`
	Next     string `json:"next"`
	Offset   int    `json:"offset"`
	Previous string `json:"previous"`
	Total    int    `json:"total"`
}

// Get returns the most recent version of the object in its entirety
func (im *PagingMeta) Get(api *API, obj interface{}) error {
	return im.HREF.Get(api, obj)
}

type AlbumPage struct {
	PagingMeta
	Items []*Album `json:"items"`
}

type TrackPage struct {
	PagingMeta
	Items []*Track `json:"items"`
}

type PlaylistPage struct {
	PagingMeta
	Items []*Playlist `json:"items"`
}

type PlaylistTrackPage struct {
	PagingMeta
	Items []*PlaylistTrack `json:"items"`
}
