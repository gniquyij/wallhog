package main

import (
    "flag"
    "time"
    "wallhog/hog"
    "wallhog/tunnels"
)

func main() {
    var tag string
    flag.StringVar(&tag, "t", "", "Specify a tag. Default is blank.")
    flag.Parse()
    keyword := flag.Args()[0]
    tunnelUrls := tunnels.GetUrls("./tunnels/tunnels.json", tag)
    for _, tunnel := range tunnelUrls {
        go hog.OpenUrl(tunnel, keyword) // if multi tags
    }
    time.Sleep(3600 * time.Second)
}