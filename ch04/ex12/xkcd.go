package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

type Comic struct {
	Id         int
	SafeTitle  string `json:"safe_title"`
	Transcript string
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("failed")
	}
	w := strings.ToLower(os.Args[1])

	files, err := ioutil.ReadDir("./data")
	if err != nil {
		log.Fatal("failed")
	}

	for _, f := range files {
		raw, err := ioutil.ReadFile(path.Join("./data", f.Name()))
		if err != nil {
			log.Fatal(err)
		}

		var comic Comic
		json.Unmarshal(raw, &comic)
		if strings.Index(comic.SafeTitle, w) > -1 {
			fmt.Printf("title:%v\thttps://xkcd.com/%v/info.0.json\nTranscript:%v", comic.SafeTitle, comic.Id, comic.Transcript)
		}
	}
}
