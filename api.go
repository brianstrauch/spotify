package spotify

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
)

const APIHost = "api.spotify.com"

// Error represents an ErrorObject in the Spotify API
// https://developer.spotify.com/documentation/web-api/reference/#object-errorobject
type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type API struct {
	token string
}

func NewAPI(token string) *API {
	return &API{token}
}

func (a *API) get(apiVersion, endpoint string, query url.Values, res interface{}) error {
	return a.call(http.MethodGet, apiVersion, endpoint, query, nil, res)
}

func (a *API) post(apiVersion, endpoint string, query url.Values, body io.Reader) error {
	return a.call(http.MethodPost, apiVersion, endpoint, query, body, nil)
}

func (a *API) put(apiVersion, endpoint string, query url.Values, body io.Reader) error {
	return a.call(http.MethodPut, apiVersion, endpoint, query, body, nil)
}

func (a *API) delete(apiVersion, endpoint string, query url.Values) error {
	return a.call(http.MethodDelete, apiVersion, endpoint, query, nil, nil)
}

func (a *API) call(method, apiVersion, endpoint string, query url.Values, body io.Reader, result interface{}) error {
	url := url.URL{
		Host:     APIHost,
		Path:     path.Join(apiVersion, endpoint),
		RawQuery: query.Encode(),
		Scheme:   "https",
	}

	req, err := http.NewRequest(method, url.String(), body)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.token))

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Success
	if res.StatusCode >= 200 && res.StatusCode < 300 {
		if result != nil {
			if err := json.NewDecoder(res.Body).Decode(result); err != nil {
				return err
			}
		}
		return nil
	}

	// Error
	spotifyErr := &struct {
		Error Error `json:"error"`
	}{}
	if err := json.NewDecoder(res.Body).Decode(spotifyErr); err != nil {
		return err
	}

	return errors.New(spotifyErr.Error.Message)
}
