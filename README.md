# Spotify Go Library

<img src="https://storage.googleapis.com/pr-newsroom-wp/1/2018/11/Spotify_Logo_RGB_Green.png" width="400">

## Description

A Go library for the Spotify Web API and Accounts service.

* Authenticate with PKCE
* Get currently playing music or podcast
* Search for specific songs by name
* Save music to user's library

## Download

```
go get github.com/brianstrauch/spotify
```

## Usage
```go
import (
  "fmt"
  "github.com/brianstrauch/spotify"
)

const token = "<YOUR API TOKEN>"

func main() {  
  api := spotify.NewAPI(token)
  
  if err := api.Play(); err != nil {
    panic(err)
  }
  
  playback, err := api.GetPlayback()
  if err != nil {
    panic(err)
  }
  
  fmt.Printf("Playing %s\n", playback.Item.Name)
}
```

## Used By
* https://github.com/brianstrauch/spotify-cli
