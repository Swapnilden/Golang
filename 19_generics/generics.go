package main

import "fmt"

func printSlice[T comparable, V string](items []T, name V) {
	for _, item := range items {
		fmt.Println(item, name)
	}
}

// func printStringSlice(items []string) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// type stack[T any] struct {
// 	element []T
// }

func main() {
	// myStack := stack[string]{
	// 	element: []string{"golang"},
	// }
	// fmt.Println(myStack)
	// nums := []int{1, 2, 3}
	// printSlice(nums)
	// names := []string{"golang", "typescript"}
	vals := []bool{true, false, true}
	printSlice(vals, "john")
	// printStringSlice(names)
}
