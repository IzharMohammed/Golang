package main

import "fmt"

func main() {
	// by default values
	// int -> 0, string -> "", bool -> false
	// var nums [4]int

	// nums[2] = 20
	// // printing array
	// fmt.Println(nums)
	// // array length
	// fmt.Println(len(nums))

	// var vals[4] bool
	// vals[2] = true
	// fmt.Println(vals)

	// to declare it in single line
	// nums:= [3]int{1,2,3}
	// fmt.Println(nums)

	//2d Arrays
	nums := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(nums)

	// -> Arrays should be used in go when there is fixed size, that is predictable
	// -> Memory optimization
	// -> constant time access
	
}
