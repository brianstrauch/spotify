package spotify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
)

// GetPlaylists gets a list of the playlists owned or followed by the current Spotify user.
// https://developer.spotify.com/documentation/web-api/reference/#endpoint-get-a-list-of-current-users-playlists
func (a *API) GetPlaylists() ([]*Playlist, error) {
	playlistPage := &struct {
		PagingMeta
		Items []*Playlist `json:"items"`
	}{}

	// TODO: Iterate over all pages of playlists

	err := a.get("v1", "/me/playlists", nil, playlistPage)
	return playlistPage.Items, err
}

// CreatePlaylist creates a playlist for a Spotify user. (The playlist will be empty until you add tracks.)
// https://developer.spotify.com/documentation/web-api/reference/#endpoint-create-playlist
func (a *API) CreatePlaylist(userID, name string, public, collaborative bool, description string) (*Playlist, error) {
	query := make(url.Values)
	query.Add("user_id", userID)

	body := &struct {
		Name          string `json:"name"`
		Public        bool   `json:"public"`
		Collaborative bool   `json:"collaborative"`
		Description   string `json:"description"`
	}{
		Name:          name,
		Public:        public,
		Collaborative: collaborative,
		Description:   description,
	}

	data, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	playlist := new(Playlist)
	err = a.post("v1", fmt.Sprintf("/users/%s/playlists", userID), query, bytes.NewReader(data), playlist)

	return playlist, err
}

// GetPlaylist gets a playlist owned by a Spotify user.
// https://developer.spotify.com/documentation/web-api/reference/#endpoint-get-playlist
func (a *API) GetPlaylist(id string) (*Playlist, error) {
	playlist := new(Playlist)
	err := a.get("v1", fmt.Sprintf("/playlists/%s", id), nil, playlist)
	return playlist, err
}
