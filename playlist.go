package spotify

import (
	"path"
)

type Playlist struct {
	Meta
	Collaborative bool
	Description   string
	Followers     ItemsMeta
	Images        []Image // TODO: what is this
	Name          string
	Owner         Owner
	Public        bool
	SnapshotID    string `json:"snapshot_id"`
	Tracks        Tracks
}

type Owner struct {
	Meta
}

type Response struct {
	ItemsMeta
	Items []Playlist
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
