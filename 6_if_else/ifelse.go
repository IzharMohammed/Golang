package main

import "fmt"

func main() {
	// age := 20

	// if age >= 18 {
	// 	fmt.Println("person is an adult")
	// }else{
	// 	fmt.Println("person is not an adult")
	// }

	// age := 20

	// if age >= 18 {
	// 	fmt.Println("person is an adult")
	// } else if age >= 12 {
	// 	fmt.Println("person is a teenager")
	// } else {
	// 	fmt.Println("person is a kid")
	// }

	// && operator

	// var role = "admin"
	// var hasPermission = false

	// if role == "admin" && hasPermission {
	// 	fmt.Println("has permission")
	// } else {
	// 	fmt.Println("No permission")
	// }

	// || operator
	// var role = "admin"
	// var hasPermission = false

	// if role == "admin" || hasPermission {
	// 	fmt.Println("has permission")
	// } else {
	// 	fmt.Println("No permission")
	// }

	// we can also declare variable inside if
	if age := 15; age >= 18 {
		fmt.Println("Person is an adult", age)
	} else if age >= 12 {
		fmt.Println("Person is a teenager", age)
	}


	// go does not have ternary operator, u will have to use normal if else

}
