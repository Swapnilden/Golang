package main

import "fmt"

func main() {
	// nums := []int{6, 7, 8}

	// for i, num := range nums {
	// 	fmt.Println(i, num)
	// }

	// m := map[string]string{"fname": "john", "lname": "doe"}
	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }

	// for k := range m {
	// 	fmt.Println(k)
	// }

	// unicode code point rune
	// starting byte of rune
	// 300 -> 1 byte, 2 byte (after 255 val 2 byte will be taken to store char)
	for i, c := range "golang" {
		fmt.Println(i, string(c))
	}
}
