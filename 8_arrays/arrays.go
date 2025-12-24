package main

import "fmt"

func main() {
	// zeroed values
	// int -> 0, string -> "", bool -> false
	// var nums [4]int

	// nums[0] = 1
	// fmt.Println(nums)
	// fmt.Println(nums[0])

	// length of array
	// fmt.Println(len(nums))

	// var vals [4]bool
	// vals[2] = true
	// fmt.Println(vals)

	// var name [3]string
	// name[0] = "golang"
	// fmt.Println(name)

	// nums := [3]int{1, 2, 3}
	// fmt.Println(nums)

	nums := [2][2]int{{3, 4}, {5, 6}}
	fmt.Println(nums)

	// - fixed size, that is predictable
	// - Memory optimization
	// - constant access time
}
