package reddit

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type RedditListing struct {
	Kind           string `json:"kind"`
	RedditMetaData struct {
		Modhash string         `json:"modhash"`
		Dist    int            `json:"dist"`
		Posts   []PostMetaData `json:"children"`
	} `json:"data"`
}

type PostMetaData struct {
	Kind string `json:"kind"`
	Post Post   `json:"data"`
}

type Post struct {
	Title string `json:"title"`
	Link  string `json:"url"`
}

func GetTopPosts() []PostMetaData {
	url := "https://www.reddit.com/r/all/.json?"
	client := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "reddit-post")

	res, getErr := client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	data := RedditListing{}
	jsonErr := json.Unmarshal(body, &data)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return data.RedditMetaData.Posts
}
