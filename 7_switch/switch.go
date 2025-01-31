package main

import (
	"fmt"
	
)

func main() {
	// simple switch
	// i := -1

	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("Three")
	// default:
	// 	fmt.Println("other")
	// }

	// Multiple condition switch
	// switch time.Now().Weekday() {
	// case time.Sunday, time.Saturday:
	// 	fmt.Println("It's weekend", time.Now().Weekday())
	// default:
	// 	fmt.Println("It's working day...", time.Now().Weekday())
	// }

	// type switch
	// interface{} means it can accept any value
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("It's an integer")
		case string:
			fmt.Println("it's an string")
		case bool:
			fmt.Println("It's an boolean")
		default:
			fmt.Println("It's the other type", t)
		}
	}

	whoAmI(true)

}
