package main

import (
	"fmt"
	"time"
)

// Channels in Go
// ----------------
// Channels are a way to communicate between goroutines safely without explicit locks.
// They help synchronize goroutines and allow safe data transfer.
//
// Benefits of channels:
// 1. Synchronization: Helps avoid race conditions.
// 2. Communication: Allows goroutines to exchange data.
// 3. Avoids manual locking: No need to use mutexes.
//
// Creating channels:
// var ch chan int          // Declaring a channel
// ch = make(chan int)      // Creating an unbuffered channel
// ch := make(chan int, 10) // Creating a buffered channel

// Function to receive data from a channel and process it
/* func processNum(numChan chan int) {
	for num := range numChan {
		fmt.Println("Processing number is...", num)
		time.Sleep(time.Second) // Simulating some processing time
	}
} */

// Function to send sum of two numbers to a channel
func sum(result chan int, num1 int, num2 int) {
	numResult := num1 + num2
	result <- numResult // Sending result to the channel
}

// Function to synchronize goroutine execution using a channel
func task(done chan bool) {
	defer func() {
		done <- true // Notify that the task is completed
	}()

	fmt.Println("Processing ....")
}

// Function to simulate an email sender that reads from a channel
func emailSender(emailChan chan string, done chan bool) {
	defer func() { done <- true }() // Notify that sending is complete
	for email := range emailChan {
		fmt.Println("Sending email to", email)
		time.Sleep(time.Second) // Simulating sending time
	}
}

func main() {
	// --- Example: Sending data through channels ---

	/*
	// Creating an unbuffered channel
	numChan := make(chan int)

	// Starting a goroutine to process numbers received from the channel
	go processNum(numChan)

	// Sending random numbers to the channel
	for {
		numChan <- rand.Intn(100)
	}
	*/

	// --- Example: Receiving data through a channel ---
	result := make(chan int) // Creating a channel to receive an integer result
	go sum(result, 5, 4)     // Starting a goroutine to calculate sum
	res := <-result          // Receiving the result from the channel
	fmt.Println(res)

	// --- Example: Goroutine Synchronization ---
	/*
	done := make(chan bool) // Creating a channel to synchronize goroutines
	go task(done)           // Starting a goroutine for a task
	<-done                  // Waiting for the task to complete
	*/

	// --- Example: Buffered Channel ---
	emailChan := make(chan string, 100) // Creating a buffered channel with capacity 100
	done := make(chan bool)             // Channel to notify when emails are sent

	// Starting a goroutine for email sending
	go emailSender(emailChan, done)

	// Sending email addresses to the channel
	for i := 0; i < 5; i++ {
		emailChan <- fmt.Sprintf("%d@gmail.com", i)
	}
	fmt.Println("Done sending emails")

	// Important: Closing the channel to indicate no more data will be sent
	close(emailChan)
	<-done // Waiting for emailSender goroutine to finish

	// --- Example: Handling Multiple Channels with select ---
	chan1 := make(chan int)    // Creating an integer channel
	chan2 := make(chan string) // Creating a string channel

	// Goroutine to send data to chan1
	go func() {
		chan1 <- 10
	}()

	// Goroutine to send data to chan2
	go func() {
		chan2 <- "pong"
	}()

	// Using select to wait on multiple channels
	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("Received data from chan1:", chan1Val)
		case chan2Val := <-chan2:
			fmt.Println("Received data from chan2:", chan2Val)
		}
	}

	// --- Example: Channel Basics (Intro to using channels) ---
	/*
	time.Sleep(time.Second * 2)

	messageChan := make(chan string) // Creating a string channel

	messageChan <- "izhar" // Sending a value into the channel (will cause deadlock if not received)

	msg := <-messageChan // Receiving from the channel

	fmt.Println(msg)
	*/
}