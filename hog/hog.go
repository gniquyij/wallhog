package hog

import "os/exec"

func main() {
    exec.Command("open", "https://www.bing.com/search?q=a").Run() // mac
}