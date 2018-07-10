package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/philippklemmer/pingRedditPost/reddit"
)

func handler(w http.ResponseWriter, r *http.Request) {
	data, err := reddit.GetTopPosts()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	t, _ := template.ParseFiles("index.html")
	t.Execute(w, data)
}

func main() {

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
