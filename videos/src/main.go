package main

import "fmt"

func main() {

	videos := getVideos()

	for i, video := range videos {
		videos[i].Description = video.Description + ". Remember to like it"
	}

	fmt.Println(videos)

}
