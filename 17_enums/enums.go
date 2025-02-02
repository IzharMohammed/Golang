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
