package reddit

import (
	"encoding/json"
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
func GetTopPosts() ([]PostMetaData, error) {
	url := "https://www.reddit.com/r/all/.json?"
	client := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "reddit-post")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	data := Listing{}
	if err = json.NewDecoder(res.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data.MetaData.Posts, nil
}
