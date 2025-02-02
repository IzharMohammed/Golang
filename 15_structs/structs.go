package main

import (
	"fmt"
	"time"
)

// Defining a struct named "customer" that holds customer details
type customer struct {
	name  string // Name of the customer
	phone string // Phone number of the customer
}

// Defining a struct named "order" that represents an order
// This struct embeds the "customer" struct, meaning it inherits its fields
type order struct {
	id        string    // Unique identifier for the order
	amount    float32   // Amount of the order
	status    string    // Status of the order (e.g., confirmed, pending)
	createdAt time.Time // Timestamp of when the order was created
	customer             // Embedded customer struct
}

// Constructor function to create a new order
func newOrder(id string, amount float32, status string) *order {
	myOrder := order{
		id:     id,
		amount: amount,
		status: status,
	}
	myOrder.createdAt = time.Now() // Set the current timestamp

	return &myOrder // Return a pointer to the new order
}

func main() {
	// Creating an anonymous struct (a struct without a named type)
	language := struct {
		name   string // Name of the programming language
		isGood bool   // Boolean flag indicating whether it's good
	}{"golang", true} // Initializing the struct inline

	// Creating an order struct using embedded struct fields
	myOrder := order{
		id:     "1",
		amount: 10.2,
		status: "confirmed",
		customer: customer{ // Initializing the embedded struct "customer"
			name:  "izhar",
			phone: "88828282",
		},
	}

	// Another way of struct initialization using a separate customer variable
	newCustomer := customer{
		name:  "izhar",
		phone: "88828282",
	}

	// Creating an order using the newCustomer instance
	myOrder2 := order{
		id:       "1",
		amount:   10.2,
		status:   "confirmed",
		customer: newCustomer, // Assigning a pre-defined customer instance
	}

	// Printing the struct instances
	fmt.Println(myOrder)  // Output the first order with embedded customer
	fmt.Println(myOrder2) // Output the second order with separately initialized customer
	fmt.Println(language) // Output the anonymous struct instance
}
