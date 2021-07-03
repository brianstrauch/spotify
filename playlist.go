package spotify

import (
	"fmt"
	"time"
)

// PlaylistTrack represents a PlaylistTrackObject in the Spotify API
// https://developer.spotify.com/documentation/web-api/reference/#object-playlisttrackobject
type PlaylistTrack struct {
	AddedAt time.Time `json:"added_at"`
	AddedBy Meta      `json:"added_by"`
	IsLocal bool      `json:"is_local"`
	Track   Track     `json:"track"`
	URI     string    `json:"uri"`
}

// Playlist represents a PlaylistObject in the Spotify API
// https://developer.spotify.com/documentation/web-api/reference/#object-playlistobject
type Playlist struct {
	Meta
	Collaborative bool              `json:"collaborative"`
	Description   string            `json:"description"`
	Images        []Image           `json:"images"`
	Name          string            `json:"name"`
	Owner         PublicUser        `json:"owner"`
	Public        bool              `json:"public"`
	SnapshotID    string            `json:"snapshot_id"`
	Tracks        PlaylistTrackPage `json:"tracks"`
}

func (a *API) GetPlaylists() ([]*Playlist, error) {
	playlists := new(PlaylistPage)
	if err := a.get("v1", "/me/playlists", nil, playlists); err != nil {
		return nil, err
	}

	return playlists.Items, nil
}

func (a *API) GetPlaylist(id string) (*Playlist, error) {
	playlist := new(Playlist)
	if err := a.get("v1", fmt.Sprintf("/playlists/%s", id), nil, playlist); err != nil {
		return nil, err
	}

	return playlist, nil
}
