package main

import (
    "wallhog/hog"
)

func main() {
    tunnels := []string{
        "https://www.bing.com/search?q=a",
        "https://www.bing.com/search?q=ab",
        "https://www.bing.com/search?q=ac",
    }
    for _, tunnel := range tunnels {
        hog.OpenUrl(tunnel)
    }
}