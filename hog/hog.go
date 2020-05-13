package hog

import (
    "os/exec"
)

func OpenUrl(url string) {
    exec.Command("open", url).Run() // mac
}