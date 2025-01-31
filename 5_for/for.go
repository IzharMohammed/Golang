package main

import "fmt"

// for --> only construct in go for looping

func main() {
	// while loop
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// infinite loop
	// for {
	// 	fmt.Println("izhar is the greatest...!!!")
	// }

	// classic for loop
	for i := 0; i < 10; i++ {

		if i == 5 {
			continue
		}
		fmt.Println(i)
	}

	//1.22 version
	for v := range 20 {
		fmt.Println(v)
	}

}
