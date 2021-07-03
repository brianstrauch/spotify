package spotify

// Meta represents common fields found in most API responses
type Meta struct {
	HREF         *HREF
	ExternalURLs map[string]string `json:"external_urls"`
	ID           string            `json:"id"`
	Type         string            `json:"type"`
	URI          string            `json:"uri"`
}

// Get returns the most recent version of the object in its entirety
func (m *Meta) Get(api *API, obj interface{}) error {
	return m.HREF.Get(api, obj)
}

// PagingMeta represents common fields in paged responses
type PagingMeta struct {
	HREF     *HREF
	Limit    int     `json:"limit"`
	Next     *string `json:"next"`
	Offset   int     `json:"offset"`
	Previous *string `json:"previous"`
	Total    int     `json:"total"`
}

// Get returns the most recent version of the object in its entirety
func (im *PagingMeta) Get(api *API, obj interface{}) error {
	return im.HREF.Get(api, obj)
}
