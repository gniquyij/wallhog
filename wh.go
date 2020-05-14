package main

import (
    "wallhog/hog"
    "time"
)

func main() {
    tunnels := []string{
        "https://www.bing.com/search?q=a",
        "https://www.bing.com/search?q=ab",
        "https://www.bing.com/search?q=ac",
    }
    for _, tunnel := range tunnels {
        go hog.OpenUrl(tunnel)
    }
    time.Sleep(3600 * time.Second)
}