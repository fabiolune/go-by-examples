package main

import (
	"fmt"
	"time"
)

func main() {
	foods := []string{"pizza", "pasta", "noodles", "poke", "steak", "fish", "bread"}

	// instantiate the channel
	results := make(chan bool)

	startTime := time.Now()

	for _, food := range foods {
		go func(f string) {
			cookFood(f)
			// fill the channel
			results <- true
		}(food)
	}

	for i := 0; i < len(foods); i++ {
		// free the channel
		<-results
	}

	elapsed := time.Since(startTime)
	fmt.Println("passed", elapsed.Milliseconds(), "ms")
}

func cookFood(food string) {
	fmt.Println("start cooking ", food)
	time.Sleep(2 * time.Second)
	fmt.Println(food, "cooked")
}
