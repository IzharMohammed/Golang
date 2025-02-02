package main

import (
	"fmt"
	"time"
)

// Goroutines allow concurrent execution of functions, making Go highly efficient in handling multiple tasks.
func task(id int) {
	fmt.Println("doing task:-", id)
}

func main() {
	for i := 0; i <= 10; i++ {
		// Using a goroutine to execute the task concurrently
		go task(i)
		
		// Anonymous function with goroutine
		// Note: Capturing loop variable correctly is important to avoid unexpected behavior
		go func(i int) {
			fmt.Println("doing task:-", i)
		}(i) // Passing i as an argument to ensure correct capture
	}

	// Adding sleep to allow goroutines to complete execution before program exits
	time.Sleep(time.Second * 2)
}

/*
Why Use Goroutines in Go?
1. **Concurrency**: Goroutines enable efficient concurrent programming by allowing multiple functions to run simultaneously.
2. **Lightweight**: Unlike traditional threads, goroutines consume minimal memory and resources.
3. **Scalability**: Allows applications to handle thousands of concurrent tasks efficiently.
4. **Performance**: Reduces execution time by running tasks in parallel.

How to Use Goroutines in Go?
- Use the `go` keyword before a function call to run it as a goroutine.
- Be cautious of loop variables when using anonymous functions inside loops.
- Use synchronization mechanisms (e.g., WaitGroups, Channels) to control goroutine execution.
- Add appropriate sleep or synchronization techniques to ensure goroutines complete execution before the program exits.
*/
