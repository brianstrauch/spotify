package spotify

import (
	"path"
)

type Playlist struct {
	Meta
	Collaborative bool      `json:"collaborative"`
	Description   string    `json:"description"`
	Followers     ItemsMeta `json:"followers"`
	Images        []Image   `json:"images"`
	Name          string    `json:"name"`
	Owner         Owner     `json:"owner"`
	Public        bool      `json:"public"`
	SnapshotID    string    `json:"snapshot_id"`
	Tracks        Tracks    `json:"tracks"`
}

type Owner struct {
	Meta
}

type Response struct {
	ItemsMeta
	Items []Playlist `json:"items"`
}

func (a *API) GetPlaylists() ([]Playlist, error) {
	playlistsResponse := new(Response)
	if err := a.get("/me/playlists", playlistsResponse); err != nil {
		return nil, err
	}
	// TODO: do more with response
	return playlistsResponse.Items, nil
}

func (a *API) GetPlaylist(playlistID string) (*Playlist, error) {
	playlistResponsse := new(Playlist)
	if err := a.get(path.Join("/playlists", playlistID), playlistResponsse); err != nil {
		return nil, err
	}
	return playlistResponsse, nil
}
