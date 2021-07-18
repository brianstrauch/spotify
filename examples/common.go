package examples

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/pkg/browser"

	"github.com/brianstrauch/spotify"
)

const (
	ClientID    = "81dddfee3e8d47d89b7902ba616f3357"
	RedirectURI = "http://localhost:1024/callback"
)

func Login() string {
	verifier, challenge, err := spotify.CreatePKCEVerifierAndChallenge()
	if err != nil {
		panic(err)
	}

	state, err := spotify.GenerateRandomState()
	if err != nil {
		panic(err)
	}

	uri := spotify.BuildPKCEAuthURI(ClientID, RedirectURI, challenge, state)

	if err := browser.OpenURL(uri); err != nil {
		panic(err)
	}

	code, err := ListenForCode(state)
	if err != nil {
		panic(err)
	}

	token, err := spotify.RequestPKCEToken(ClientID, code, RedirectURI, verifier)
	if err != nil {
		panic(err)
	}

	return token.AccessToken
}

func ListenForCode(state string) (code string, err error) {
	server := &http.Server{Addr: ":1024"}

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

	return
}
