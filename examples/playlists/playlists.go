package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/brianstrauch/spotify"
	"github.com/pkg/browser"
)

const (
	clientID    = "81dddfee3e8d47d89b7902ba616f3357"
	redirectURI = "http://localhost:1024/callback"
)

func main() {
	verifier, challenge, err := spotify.CreateVerifierAndChallenge()
	if err != nil {
		panic(err)
	}

	state, err := spotify.GenerateRandomState()
	if err != nil {
		panic(err)
	}

	uri := spotify.BuildAuthURI(clientID, redirectURI, challenge, state)

	if err := browser.OpenURL(uri); err != nil {
		panic(err)
	}

	code, err := listenForCode(state)
	if err != nil {
		panic(err)
	}

	token, err := spotify.RequestToken(clientID, code, redirectURI, verifier)
	if err != nil {
		panic(err)
	}

	playlists, err := spotify.NewAPI(token.AccessToken).GetPlaylists()
	if err != nil {
		panic(err)
	} else if len(playlists) == 0 {
		panic("no playlists")
	}

	playlist := new(spotify.Playlist)
	if err := playlists[0].Get(spotify.NewAPI(token.AccessToken), playlist); err != nil {
		panic(err)
	}
	fmt.Println(playlist.Name)
	fmt.Printf("%d tracks\n", playlist.Tracks.Total)
}

func listenForCode(state string) (string, error) {
	server := &http.Server{Addr: ":1024"}

	var code string
	var err error

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("state") != state || r.URL.Query().Get("error") != "" {
			err = errors.New("authorization failed")
			fmt.Fprintln(w, "Failure.")
		} else {
			code = r.URL.Query().Get("code")
			fmt.Fprintln(w, "Success!")
		}

		// Use a separate thread so browser doesn't show a "No Connection" message
		go func() {
			server.Shutdown(context.Background())
		}()
	})

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		return "", err
	}

	return code, err
}
