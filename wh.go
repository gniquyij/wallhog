package main

import (
    "flag"
    "os"
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
    whPath := os.Getenv("WALLHOG_PATH")
    dataPath := filepath.Join(whPath, "/tunnels/tunnels.json")
    tunnelUrls := tunnels.GetUrls(dataPath, tag)
    for _, tunnel := range tunnelUrls {
        go hog.OpenUrl(tunnel, keyword) // if multi tags
    }
    time.Sleep(3600 * time.Second)
}