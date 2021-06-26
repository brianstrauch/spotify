package spotify

import (
	"net/url"
	"strings"
)

type HREF string

func (h *HREF) Get(api *API, obj interface{}) error {
	url, err := h.GetURL()
	if err != nil {
		return err
	}
	path := strings.Replace(url.Path, "/v1", "", 1)
	return api.get(path, obj)
}

func (h *HREF) GetURL() (*url.URL, error) {
	return url.Parse(string(*h))
}

type Meta struct {
	HREF         *HREF
	ExternalURLs map[string]string `json:"external_urls"`
	ID           string            `json:"id"`
	Type         string            `json:"type"`
	URI          string            `json:"uri"`
}

func (m *Meta) Get(api *API, obj interface{}) error {
	return m.HREF.Get(api, obj)
}

type ItemsMeta struct {
	HREF     *HREF
	Limit    int     `json:"limit"`
	Next     *string `json:"next"`
	Offset   int     `json:"offset"`
	Previous *string `json:"previous"`
	Total    int     `json:"total"`
}

func (im *ItemsMeta) Get(api *API, obj interface{}) error {
	return im.HREF.Get(api, obj)
}
