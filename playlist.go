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

// GetPlaylists gets a list of the playlists owned or followed by the current Spotify user.
// https://developer.spotify.com/documentation/web-api/reference/#endpoint-get-a-list-of-current-users-playlists
func (a *API) GetPlaylists() ([]*Playlist, error) {
	playlists := new(PlaylistPage)
	err := a.get("v1", "/me/playlists", nil, playlists)
	return playlists.Items, err
}

// GetPlaylist gets a playlist owned by a Spotify user.
// https://developer.spotify.com/documentation/web-api/reference/#endpoint-get-playlist
func (a *API) GetPlaylist(id string) (*Playlist, error) {
	playlist := new(Playlist)
	err := a.get("v1", fmt.Sprintf("/playlists/%s", id), nil, playlist)
	return playlist, err
}
