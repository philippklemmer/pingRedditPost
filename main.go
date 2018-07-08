package main

import (
	"html/template"
	"log"
	"net/http"
	"pingRedditPost/reddit"
)

func handler(w http.ResponseWriter, r *http.Request) {
	data := reddit.GetTopPosts()
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, data)
}

func main() {

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
