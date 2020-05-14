package tunnels

import (
    "wallhog/tunnels/bing"
    "wallhog/tunnels/ecosia"
)

var Tunnels = []string{
    ecosia.Ecosia.Url,
    bing.Bing.Url,
}