package tunnels

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
)

type Tunnel struct {
    Name, Url string
}

type Tunnels []Tunnel

func GetTunnelData(file string) Tunnels {
    data, err := ioutil.ReadFile(file)
    if err != nil {
        fmt.Print(err)
    }

    var TunnelData Tunnels
    err = json.Unmarshal(data, &TunnelData)
    if err != nil {
        fmt.Println("error:", err)
    }

    return TunnelData
}

func GetTunnelUrls(file string) []string {
    var TunnelData = GetTunnelData(file)

    var TunnelUrls = []string{}
    for i := 0; i < len(TunnelData); i++ {
        TunnelUrls = append(TunnelUrls, TunnelData[i].Url)
    }

    return TunnelUrls
}