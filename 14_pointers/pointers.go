package main

import "fmt"

// func changeNum(num int) {
// 	num = 5
// 	fmt.Println("In changeNum", num)
// }

func changeNum(num *int) {
	*num = 5
	fmt.Println("In change num", *num)
}

func main() {
	num := 1
	// changeNum(num)
	changeNum(&num)
	fmt.Println("After changing num in main", num)
}
