package main

import "fmt"

// func add(a int, b int) int {
// 	return a + b
// }

// func add(a, b int) int {
// 	return a + b
// }

// func getLanguages() (string, string, string) {
// 	return "golang", "js", "rust"
// }

// passing function as argument
// func processIt(fn func(a int) int) {
// 	fn(1)
// }

// returning function as value
func processIt() func(a int) int {
	return func(a int) int {
		return a
	}
}

func main() {
	// lang1, _, lang3 := getLanguages()
	// fmt.Println(lang1, lang3)

	// fmt.Println(add(1, 2))

	// fn := func(a int) int {
	// 	return 2
	// }
	// processIt(fn)

	fn := processIt()
	fmt.Println(fn(6))
}
