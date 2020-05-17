package main

import (
    "flag"
    "os"
    "path/filepath"
    "time"
    "wallhog/hog"
    "wallhog/tunnels"
)

func main() {
    var tag string
    flag.StringVar(&tag, "t", "general", "Specify a tag. Default is general.")
    flag.Parse()
    keyword := flag.Args()[0]
//     currentPath, _ := os.Getwd()
    whPath := os.Getenv("WALLHOG_PATH")
//     whPath := "/Users/yuqing.ji/wallhog"
    dataPath := filepath.Join(whPath, "/tunnels/tunnels.json")
    tunnelUrls := tunnels.GetUrls(dataPath, tag)
    for _, tunnel := range tunnelUrls {
        go hog.OpenUrl(tunnel, keyword) // if multi tags
    }
    time.Sleep(3600 * time.Second)
}