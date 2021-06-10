# Spotify Go Library

<img src="https://storage.googleapis.com/pr-newsroom-wp/1/2018/11/Spotify_Logo_RGB_Green.png" width="400">

## Description

A Go library for the Spotify Web API and Accounts service.

⚠️ Warning: Work in Progress ⚠️

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
  
  if err := api.Play(""); err != nil {
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
