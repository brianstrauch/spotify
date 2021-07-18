package main

import (
	"fmt"

	"github.com/pkg/browser"

	"github.com/brianstrauch/spotify"
	"github.com/brianstrauch/spotify/examples"
)

const (
	clientID     = "REDACTED"
	clientSecret = "REDACTED"
)

func main() {
	// 1. Have your application request authorization; the user logs in and authorizes access
	state, err := spotify.GenerateRandomState()
	if err != nil {
		panic(err)
	}

	uri := spotify.BuildAuthURI(clientID, examples.RedirectURI, state, false)

	if err := browser.OpenURL(uri); err != nil {
		panic(err)
	}

	code, err := examples.ListenForCode(state)
	if err != nil {
		panic(err)
	}

	// 2. Have your application request refresh and access tokens; Spotify returns access and refresh tokens
	token, err := spotify.RequestToken(clientID, code, examples.RedirectURI, clientSecret)
	if err != nil {
		panic(err)
	}
	fmt.Println(token)

	// 3. Use the access token to access the Spotify Web API; Spotify returns requested data
	user, err := spotify.NewAPI(token.AccessToken).GetUserProfile()
	if err != nil {
		panic(err)
	}
	fmt.Println(user)

	// 4. Requesting a refreshed access token; Spotify returns a new access token to your app
	token, err = spotify.RefreshToken(token.RefreshToken, clientID, clientSecret)
	if err != nil {
		panic(err)
	}
	fmt.Println(token)
}
