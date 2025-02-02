package main

import (
	"fmt"
	"sync"
)

func task(id int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("doing task:-", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i <= 10; i++ {
		wg.Add(1)
		// Using a goroutine to execute the task concurrently
		go task(i, &wg)

	}

	wg.Wait()

	// Adding sleep to allow goroutines to complete execution before program exits
	// time.Sleep(time.Second * 2)
}
