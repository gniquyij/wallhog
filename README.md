# wallhog

![demo](./demo-general.gif)

Web search from cmd. 

`$ wh <keyword>` will:
- open your default browser
- search <keyword> via the 'general' search engine(s) - search engine(s) tagged as 'general' in `./tunnels/tunnels.json`

If a specific search engine required, you could `$ wh -t <tag> <keyword>`
- E.g. `$ wh -t book <keyword>` will search <keyword> via the search engine(s) tagged as 'book' in `./tunnels/tunnels.json`

If multiple tags, you could `$ wh -t "<tag1> <tag2> ... <tagN>" <keyword>`. wallhog would search with the engines with all the tags.

## Setup

```
$ go get github.com/vjyq/wallhog
$ cd $(go env GOPATH)/src/github.com/vjyq/wallhog/bin
$ bash shortcut.sh
```
Then just restart your terminal.

## Author

yuqing.ji@outlook.com