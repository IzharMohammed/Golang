package main

import (
	"fmt"
	"sync"
)

// Mutex in Go
// ------------
// Mutex (short for "mutual exclusion") is used to prevent race conditions
// in concurrent programming by ensuring that only one goroutine can access
// a shared resource at a time.
//
// Why use Mutex?
// - Helps in synchronizing access to shared variables.
// - Prevents data races when multiple goroutines access the same resource.
// - Allows safe modifications of shared resources.
//
// How to use Mutex in Go:
// 1. Declare a `sync.Mutex` in the struct.
// 2. Use `Lock()` before modifying a shared resource.
// 3. Use `Unlock()` after modification to release the lock.
//
// Example:
// var mu sync.Mutex
// mu.Lock()
// // critical section
// mu.Unlock()

// Defining a struct `post` with a `views` counter and a Mutex
// The Mutex ensures safe concurrent access to the `views` field.
type post struct {
	views int
	mu    sync.Mutex // Mutex to prevent race conditions
}

// Function to increment the `views` counter safely
func (p *post) inc(wg *sync.WaitGroup) {
	defer func() {
		p.mu.Unlock() // Unlock after updating views
		wg.Done()     // Mark this goroutine as done
	}()

	p.mu.Lock()   // Lock before modifying shared resource
	p.views += 1 // Increment views safely
}

func main() {
	var wg sync.WaitGroup // WaitGroup to synchronize goroutines

	mypost := post{views: 0} // Creating a post instance with initial views = 0

	// Spawning 100 goroutines to increment the views counter
	for i := 0; i < 100; i++ {
		wg.Add(1) // Increment WaitGroup counter
		go mypost.inc(&wg)
	}

	wg.Wait() // Wait for all goroutines to finish execution
	fmt.Println(mypost.views) // Print the final views count
}
