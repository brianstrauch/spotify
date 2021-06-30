package main

import (
	"fmt"

	"github.com/brianstrauch/spotify"
	"github.com/brianstrauch/spotify/examples"
	"github.com/pkg/browser"
)

const (
	clientID    = "81dddfee3e8d47d89b7902ba616f3357"
	redirectURI = "http://localhost:1024/callback"
)

func main() {
	verifier, challenge, err := spotify.CreatePKCEVerifierAndChallenge()
	if err != nil {
		panic(err)
	}

	state, err := spotify.GenerateRandomState()
	if err != nil {
		panic(err)
	}

	uri := spotify.BuildPKCEAuthURI(clientID, redirectURI, challenge, state)

	if err := browser.OpenURL(uri); err != nil {
		panic(err)
	}

	code, err := examples.ListenForCode(state)
	if err != nil {
		panic(err)
	}

	token, err := spotify.RequestPKCEToken(clientID, code, redirectURI, verifier)
	if err != nil {
		panic(err)
	}

	api := spotify.NewAPI(token.AccessToken)

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
