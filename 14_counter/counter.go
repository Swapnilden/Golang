package main

import "fmt"

func counter() func() int {
	var count int = 0

	return func() int { // variable used in upper scope will always be available evenif function completes it execuion
		count += 1
		return count
	}
}

func main() {
	increment := counter()
	fmt.Println(increment())
	fmt.Println(increment())
}
