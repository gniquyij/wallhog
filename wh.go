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
    var tags string
    flag.StringVar(&tags, "t", "general", "Specify a tag. Default is general.")
    flag.Parse()
    keyword := flag.Args()[0]
    WHPATH := filepath.Join(build.Default.GOPATH, "/src/github.com/vjyq/wallhog")
    dataPath := filepath.Join(WHPATH, "/tunnels/tunnels.json")
    tunnelUrls := tunnels.GetUrls(dataPath, tags)
    for _, tunnel := range tunnelUrls {
        go hog.OpenUrl(tunnel, keyword)
    }
    time.Sleep(300 * time.Second)
}