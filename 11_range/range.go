package main

import "fmt"

//iterating over data structure
func main() {
	// For iterating over  slices
	// nums := []int{6, 7, 8}

	// for _, num := range nums {
	// 	fmt.Println(num)
	// }

	// sum := 0
	// for index, num := range nums {
	// 	fmt.Println(index)
	// 	sum = sum + num
	// }
	// fmt.Println(sum)

	// for iterating over maps
	// m := map[string]string{"name": "izhar", "role": "full stack web devOps dev"}
	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }

	//unicode code point rune
	for i, c := range "golang" {
		fmt.Println(i, c, string(c))
	}

}
