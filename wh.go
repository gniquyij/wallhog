package main

import (
    "flag"
    "go/build"
    "path/filepath"
    "time"
    "github.com/vjyq/wallhog/hog"
    "github.com/vjyq/wallhog/tunnels"
)

func main() {
    var tag string
    flag.StringVar(&tag, "t", "general", "Specify a tag. Default is general.")
    flag.Parse()
    keyword := flag.Args()[0]
    WHPATH := filepath.Join(build.Default.GOPATH, "/src/github.com/vjyq/wallhog")
    dataPath := filepath.Join(WHPATH, "/tunnels/tunnels.json")
    tunnelUrls := tunnels.GetUrls(dataPath, tag)
    for _, tunnel := range tunnelUrls {
        go hog.OpenUrl(tunnel, keyword) // if multi tags
    }
    time.Sleep(3600 * time.Second)
}