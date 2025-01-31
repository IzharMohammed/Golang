package main

import "fmt"

// we can assign constant variables outside function but we can't assign it using shorthand
const name = "golang"

// name:="golang" // syntax error: non-declaration statement outside function body

func main() {
	// const name string = "izhar"
	// name = "golang"   // error
	fmt.Println(name)

	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(port, host)
}
