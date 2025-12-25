package main

import (
	"fmt"
	"maps"
)

func main() {
	// m := make(map[string]string)

	// m["name"] = "golang"
	// m["area"] = "backend"

	// fmt.Println(m["name"], m["area"])
	// get an element
	// IMP: if a key does not exist in the map then it return zeroed value

	// m := make(map[string]int)
	// m["age"] = 20
	// m["price"] = 50
	// fmt.Println(m["age"])
	// fmt.Println(m["phone"])
	// fmt.Println(len(m))

	// delete(m, "price")

	// clear(m)
	// fmt.Println(m)

	// m := map[string]int{"price": 40, "phones": 3}
	// // fmt.Println(m)
	// v, ok := m["phones"]
	// fmt.Println(v) // if key is not present in map then it will return zeroed value
	// if ok {
	// 	fmt.Println("all ok")
	// } else {
	// 	fmt.Println("all not ok")
	// }

	m1 := map[string]int{"price": 40, "phones": 3}
	m2 := map[string]int{"price": 40, "phones": 8}

	fmt.Println(maps.Equal(m1, m2))
}
