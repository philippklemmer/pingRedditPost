package reddit

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Listing consists the initial reddit post object
type Listing struct {
	Kind     string `json:"kind"`
	MetaData struct {
		Modhash string         `json:"modhash"`
		Dist    int            `json:"dist"`
		Posts   []PostMetaData `json:"children"`
	} `json:"data"`
}

// PostMetaData consists of specific meta data and posts
type PostMetaData struct {
	Kind string `json:"kind"`
	Post struct {
		Title string `json:"title"`
		Link  string `json:"url"`
	} `json:"data"`
}

// GetTopPosts return the 20 top posts of /r/all
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

	data := Listing{}
	jsonErr := json.Unmarshal(body, &data)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return data.MetaData.Posts
}
