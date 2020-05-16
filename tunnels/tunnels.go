package tunnels

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
)

type Tunnel struct {
    Name, Url string
    Tags []string
}

type Tunnels []Tunnel

func GetAll(file string) Tunnels {
    data, err := ioutil.ReadFile(file)
    if err != nil {
        fmt.Print(err)
    }
    var All Tunnels
    err = json.Unmarshal(data, &All)
    if err != nil {
        fmt.Println("error:", err)
    }
    return All
}

func GetUrls(file, tag string) []string {
    var All = GetAll(file)
    var Urls = []string{}
    for i := 0; i < len(All); i++ {
        if tag != "" {
            tags := All[i].Tags
            set := make(map[string] bool)
            for _, v := range tags {
                set[v] = true
            }
            if set[tag] {
                Urls = append(Urls, All[i].Url)
            }
        }else{
            Urls = append(Urls, All[i].Url)
        }
    }
    return Urls
}