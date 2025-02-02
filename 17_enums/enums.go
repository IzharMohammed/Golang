package main

import "fmt"

// For int types

// type OrderStatus int

// const (
// 	Received OrderStatus = iota
// 	Confirmed
// 	Prepared
// 	Delivered
// )

// For string types

type OrderStatus string

const (
	Received  OrderStatus = "received"
	Confirmed             = "confirmed"
	Prepared              = "prepared"
	Delivered             = "delivered"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("changing order status to", status)
}

func main() {
	changeOrderStatus(Delivered)
}

/*
Why Use Enums in Go?
1. **Code Readability**: Enums provide named constants, making code easier to read and maintain.
2. **Type Safety**: Prevents invalid values by restricting them to predefined constants.
3. **Avoiding Magic Strings/Numbers**: Using meaningful names instead of raw values improves clarity.
4. **Better Maintainability**: Easier to update and extend in the future.

How to Use Enums in Go?
- Use `const` to define a set of named constants.
- Use `iota` (for int-based enums) or string literals (for string-based enums).
- Use a custom type (e.g., `type OrderStatus string`) to enforce type safety.
- Pass enums as function parameters to restrict input to valid values.
*/
