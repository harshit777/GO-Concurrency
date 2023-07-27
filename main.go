package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	// defer waits for the intial function to be completed and then run. Here it prints s first and then do wg.Done
	defer wg.Done()
	fmt.Println(s)
}

func main() {
	// As we added go routine to first function it started its own thread at the same time the main function was getting executed.
	//go printSomething("Hello World 1")

	// Wait group used for telling main func to wait till go routine is completed
	var wg sync.WaitGroup
	words := []string{
		"alpha",
		"beta",
		"gamma",
		"theta",
		"eta",
		"delta",
		"epsilon",
		"pie",
	}
	// Add  tells the delta enteries or number of executions to take place
	wg.Add(len(words))

	for i, x := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, x), &wg)
	}

	// Waits till the wg counter is zero
	wg.Wait()

	wg.Add(1)
	printSomething("Hello World 2", &wg)

	//Output
	// Helloworld 2
}
