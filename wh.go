package main

import (
    "wallhog/tunnels"
    "wallhog/hog"
    "time"
    "fmt"
)

func main() {
    for _, tunnel := range tunnels.Tunnels {
        fmt.Printf(tunnel)
        go hog.OpenUrl(tunnel)
    }
    time.Sleep(3600 * time.Second)
}