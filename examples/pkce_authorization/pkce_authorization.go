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
	// 1. Create the code verifier and challenge
	verifier, challenge, err := spotify.CreateVerifierAndChallenge()
	if err != nil {
		panic(err)
	}

	// 2. Construct the authorization URI
	state, err := spotify.GenerateRandomState()
	if err != nil {
		panic(err)
	}

	uri := spotify.BuildAuthURI(clientID, redirectURI, challenge, state)

	// 3. Your app redirects the user to the authorization URI
	if err := browser.OpenURL(uri); err != nil {
		panic(err)
	}

	code, err := listenForCode(state)
	if err != nil {
		panic(err)
	}

	// 4. Your app exchanges the code for an access token
	token, err := spotify.RequestToken(clientID, code, redirectURI, verifier)
	if err != nil {
		panic(err)
	}
	fmt.Println(token.AccessToken)

	// 5. Use the access token to access the Spotify Web API
	paging, err := spotify.NewAPI(token.AccessToken).Search("Mr. Brightside", 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(paging.Tracks.Items[0].URI)

	// 6. Fetch a playlist
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

	// 6. Requesting a refreshed access token
	token, err = spotify.RefreshToken(token.RefreshToken, clientID)
	if err != nil {
		panic(err)
	}
	fmt.Println(token.AccessToken)
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
