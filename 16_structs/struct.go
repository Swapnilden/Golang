package main

import (
	"fmt"
	"time"
)

type customer struct {
	name  string
	phone string
}

// composition
type order struct {
	id         string
	amount     float32
	status     string
	Created_At time.Time // nanosecond precision
	customer
}

// reciever type - to attach function with struct instance
// func (o *order) changeStatus(status string) { // if we are changing the struct then passing by value is necessary to modify the original/actual struct
// 	o.status = status // dereferencing thing is done by struct automatically
// }

// func (o order) getAmount() float32 { // if getting something not manipulating values of struct then passing by value can be ok
// 	return o.amount
// }

// func newOrder(id string, amount float32, status string) *order {
// 	// initial setup goes here...
// 	myOrder := order{
// 		id:     id,
// 		amount: amount,
// 		status: status,
// 	}

// 	return &myOrder
// }

func main() {

	myOrder := order{
		id:     "1",
		amount: 30,
		status: "recieved",
		customer: customer{
			name:  "john",
			phone: "1234567890",
		},
	}

	myOrder.customer.name = "robin"
	fmt.Println(myOrder)

	// language := struct {
	// 	name   string
	// 	isGood bool
	// }{"golang", true}

	// fmt.Println(language)

	// if you don't set any field then default value is zero value
	// int => 0, float => 0, string => "", bool => false
	// myOrder := order{
	// 	id:     "1",
	// 	amount: 50.00,
	// 	status: "recieved",
	// }

	// myOrder.changeStatus("confirmed")
	// fmt.Println(myOrder.amount)

	// myOrder.Created_At = time.Now()

	// fmt.Println(myOrder.status)

	// myOrder2 := order{
	// 	id:         "2",
	// 	amount:     100,
	// 	status:     "delivered",
	// 	Created_At: time.Now(),
	// }

	// myOrder.status = "paid"
	// fmt.Println("Order struct", myOrder2)
	// fmt.Println("Order struct", myOrder)
}
