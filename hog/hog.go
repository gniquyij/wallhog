package hog

import (
    "os/exec"
)

func OpenUrl(tunnel, keyword string) {
    url := tunnel + keyword
    exec.Command("open", url).Run() // mac
}