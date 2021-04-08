package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"sync"
)

type Task struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	var t Task
	wg := sync.WaitGroup{}

	sem := make(chan bool, 10)
	for i := 0; i < 100; i++ {
		fmt.Printf("%v currently running goroutine(s)\n", runtime.NumGoroutine())

		wg.Add(1)

		// adds a value to the buffered channel: being a blocking operation, it will not execute untile one of the runnint routines
		// with its 'defer' call frees one place in the channel
		sem <- true
		go func(index int) {
			// at the end of the go routine say to the wg that it completed + free one place in the semaphore
			defer wg.Done()
			defer func() {
				<-sem
			}()

			res, err := http.Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/todos/%d", index))
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			if err := json.NewDecoder(res.Body).Decode(&t); err != nil {
				log.Fatal(err)
			}
			fmt.Println(t.Title)
		}(i)
	}
	wg.Wait()
}
