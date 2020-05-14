package main

import (
    "os"
    "time"
    "wallhog/hog"
    "wallhog/tunnels"
)

func main() {
    keyword := os.Args[1]
    for _, tunnel := range tunnels.Tunnels {
        go hog.OpenUrl(tunnel, keyword)
    }
    time.Sleep(3600 * time.Second)
}