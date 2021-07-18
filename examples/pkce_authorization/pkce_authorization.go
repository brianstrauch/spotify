package main

import (
	"fmt"

	"github.com/pkg/browser"

	"github.com/brianstrauch/spotify"
	"github.com/brianstrauch/spotify/examples"
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

	uri := spotify.BuildPKCEAuthURI(examples.ClientID, examples.RedirectURI, challenge, state)

	// 3. Your app redirects the user to the authorization URI
	if err := browser.OpenURL(uri); err != nil {
		panic(err)
	}

	code, err := examples.ListenForCode(state)
	if err != nil {
		panic(err)
	}

	// 4. Your app exchanges the code for an access token
	token, err := spotify.RequestPKCEToken(examples.ClientID, code, examples.RedirectURI, verifier)
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
	token, err = spotify.RefreshPKCEToken(token.RefreshToken, examples.ClientID)
	if err != nil {
		panic(err)
	}
	fmt.Println(token)
}
