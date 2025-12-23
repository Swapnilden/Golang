package main

import "fmt"

func main() {
	// age := 16

	// if age >= 18 {
	// 	fmt.Println("person is an adult")
	// } else {
	// 	fmt.Println("person is not an adult")
	// }

	// age := 10
	// if age >= 18 {
	// 	fmt.Println("person is an adult")
	// } else if age >= 12 {
	// 	fmt.Println("person is teenager")
	// } else {
	// 	fmt.Println("person is a kid")
	// }

	// var role = "admin"
	// var hasPermission = true

	// if role == "admin" && hasPermission {
	// 	fmt.Println("yes")
	// }

	// we can declare variable inside if construct.
	if age := 20; age >= 18 {
		fmt.Println("person is an adult")
	} else if age >= 12 {
		fmt.Println("person is teenager")
	}

	// go does not have ternary operator. you will have to use if else
}
