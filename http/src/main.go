package main

import (
	"encoding/json"
	"net/http"
)

func main() {

	http.HandleFunc("/", hello)
	http.HandleFunc("/videos", getVideos)
	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("test-header", "test-value")
	w.Write([]byte("home page"))
}

func getVideos(w http.ResponseWriter, r *http.Request) {

	videos := retrieveVideos()
	data := videodata{Videos: videos}

	videoBytes, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}

	w.Header().Add("content-type", "application/json")
	w.Write(videoBytes)
}

type videodata struct {
	Videos []video `json:"videos"`
}
