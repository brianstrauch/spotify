package main

import (
	"fmt"

	"github.com/brianstrauch/spotify"
	"github.com/brianstrauch/spotify/examples"
)

func main() {
	token := examples.Login()

	api := spotify.NewAPI(token)

	playlists, err := api.GetPlaylists()
	if err != nil {
		panic(err)
	}
	if len(playlists) == 0 {
		panic("no playlists found")
	}

	playlist := new(spotify.Playlist)
	if err := playlists[0].Get(api, playlist); err != nil {
		panic(err)
	}

	fmt.Println(playlist.Name)
	for i, playlistTrack := range playlist.Tracks.Items {
		fmt.Printf("%d. %s\n", i+1, playlistTrack.Track.Name)
	}
}
