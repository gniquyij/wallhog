package main

import "os/exec"

func hog(url string) {
    exec.Command("open", url).Run() // mac
}

func main() {
    tunnels := []string{
        "https://www.bing.com/search?q=a",
        "https://www.bing.com/search?q=ab",
        "https://www.bing.com/search?q=ac",
    }
    for _, tunnel := range tunnels {
        hog(tunnel)
    }
}