package main

import "fmt"

// parameters passed by value - copy of variable
// func changeNum(num int) {
// 	num = 5
// 	fmt.Println("In changeNum", num)
// }

func changeNum(num *int) {
	*num = 5
	fmt.Println("In changeNum", *num)
}

func main() {
	num := 1

	// fmt.Println("Memory Address", &num)
	// changeNum(num)
	changeNum(&num)

	fmt.Println("After changeNum in main", num)
}
