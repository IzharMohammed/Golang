package main

import "fmt"

// When recieving mutiple arguments as parameters
func add(nums ...int) int {
	total := 0

	for _, num := range nums {
		total = total + num
	}

	return total
}

func main() {
	// fmt.Println(add(1, 2, 3, 4, 5, 6, 7, 8, 9))

	// passing slices
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := add(nums...)
	fmt.Println(result)

}
