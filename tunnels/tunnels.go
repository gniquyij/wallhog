package tunnels

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "strings"
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

func GetUrls(file, tags string) []string {
    var All = GetAll(file)
    var Urls = []string{}
    var tagSlice = strings.Split(tags, " ")
    for i := 0; i < len(All); i++ {
        if tags != "" {
            tunnelTags := All[i].Tags
            set := make(map[string] bool)
            for _, v := range tunnelTags {
                set[v] = true
            }
            var tagsExisted = set[tagSlice[0]]
            for n := 1; n < len(tagSlice); n++ {
                tagsExisted = tagsExisted && set[tagSlice[n]]
            }
            if tagsExisted {
                Urls = append(Urls, All[i].Url)
            }
        }else{
            Urls = append(Urls, All[i].Url)
        }
    }
    return Urls
}