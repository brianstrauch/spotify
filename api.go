package spotify

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const BaseURL = "https://api.spotify.com/v1"

type Error struct {
	Error struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Reason  string `json:"reason"`
	} `json:"error"`
}

type API struct {
	token string
}

func NewAPI(token string) *API {
	return &API{token}
}

func (a *API) get(endpoint string, result interface{}) error {
	return a.call(http.MethodGet, endpoint, nil, result)
}

func (a *API) post(endpoint string, body io.Reader) error {
	return a.call(http.MethodPost, endpoint, body, nil)
}

func (a *API) put(endpoint string, body io.Reader) error {
	return a.call(http.MethodPut, endpoint, body, nil)
}

func (a *API) delete(endpoint string) error {
	return a.call(http.MethodDelete, endpoint, nil, nil)
}

func (a *API) call(method string, endpoint string, body io.Reader, result interface{}) error {
	req, err := http.NewRequest(method, BaseURL+endpoint, body)
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
	spotifyErr := new(Error)
	if err := json.NewDecoder(res.Body).Decode(spotifyErr); err != nil {
		return err
	}

	return errors.New(spotifyErr.Error.Message)
}
