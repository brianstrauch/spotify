package main

import (
	"fmt"

	"github.com/pkg/browser"

	"github.com/brianstrauch/spotify"
	"github.com/brianstrauch/spotify/examples"
)

const (
	clientID    = "81dddfee3e8d47d89b7902ba616f3357"
	redirectURI = "http://localhost:1024/callback"
)

func main() {
	// 1. Create the code verifier and challenge
	verifier, challenge, err := spotify.CreatePKCEVerifierAndChallenge()
	if err != nil {
		panic(err)
	}

	// 2. Construct the authorization URI
	state, err := spotify.GenerateRandomState()
	if err != nil {
		panic(err)
	}

	uri := spotify.BuildPKCEAuthURI(clientID, redirectURI, challenge, state)

	// 3. Your app redirects the user to the authorization URI
	if err := browser.OpenURL(uri); err != nil {
		panic(err)
	}

	code, err := examples.ListenForCode(state)
	if err != nil {
		panic(err)
	}

	// 4. Your app exchanges the code for an access token
	token, err := spotify.RequestPKCEToken(clientID, code, redirectURI, verifier)
	if err != nil {
		panic(err)
	}
	fmt.Println(token)

	// 5. Use the access token to access the Spotify Web API
	user, err := spotify.NewAPI(token.AccessToken).GetUserProfile()
	if err != nil {
		panic(err)
	}
	fmt.Println(user)

	// 6. Requesting a refreshed access token
	token, err = spotify.RefreshPKCEToken(token.RefreshToken, clientID)
	if err != nil {
		panic(err)
	}
	fmt.Println(token)
}
