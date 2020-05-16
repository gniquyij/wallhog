package main

import (
    "os"
    "time"
    "wallhog/hog"
    "wallhog/tunnels"
)

func main() {
    keyword := os.Args[1]
    tunnelUrls := tunnels.GetTunnelUrls("./tunnels/tunnels.json")
    for _, tunnel := range tunnelUrls {
        go hog.OpenUrl(tunnel, keyword)
    }
    time.Sleep(3600 * time.Second)
}