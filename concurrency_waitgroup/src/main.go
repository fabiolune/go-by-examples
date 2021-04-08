package main

import (
	"fmt"
	"sync"
	"time"
)

// awaitgroup is used to store references to a certain amount of routines (go func...)
// every completed routine calls the Done method, but the waitgroup shoud be put in wait mode
func main() {
	foods := []string{"pizza", "pasta", "noodles", "poke"}

	var wg sync.WaitGroup
	wg.Add(len(foods))

	now := time.Now()
	for _, food := range foods {
		go func(f string) {
			cookFood(f)
			wg.Done()
		}(food)
	}
	wg.Wait()
	elapsed := time.Since(now)
	fmt.Println("passed", elapsed.Milliseconds(), "ms")
}

func cookFood(food string) {
	fmt.Println("start cooking ", food)
	time.Sleep(2 * time.Second)
	fmt.Println(food, "cooked")
}
