package spotify

import (
	"net/url"
	"strings"
)

type HREF struct {
	*url.URL
}

func (h *HREF) Get(api *API, obj interface{}) error {
	path := strings.Replace(h.Path, "/v1", "", 1)
	return api.get(path, obj)
}

func (h *HREF) UnmarshalJSON(data []byte) error {
	u, err := url.Parse(string(data[1 : len(data)-1]))
	if err != nil {
		return err
	}
	h.URL = u
	return nil
}

type Meta struct {
	HREF         *HREF
	ExternalURLs map[string]string `json:"external_urls"`
	ID           string
	Type         string
	URI          string
}

func (m *Meta) Get(api *API, obj interface{}) error {
	return m.HREF.Get(api, obj)
}

type ItemsMeta struct {
	HREF     *HREF
	Limit    int
	Next     *string
	Offset   int
	Previous *string
	Total    int
}

func (im *ItemsMeta) Get(api *API, obj interface{}) error {
	return im.HREF.Get(api, obj)
}
