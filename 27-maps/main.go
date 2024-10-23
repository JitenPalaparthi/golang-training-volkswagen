package main

import "fmt"

func main() {

	var map1 map[string]string // declare a map

	if map1 == nil {
		map1 = make(map[string]string, 10)
		println("map1 was nil and instantiated")
	}

	map1["522001"] = "Guntur-1"
	map1["560086"] = "Yeshvantpur-Bangalore"
	map1["560090"] = "Rajajinagar-Bangalore"

	for k, v := range map1 {
		println("key:", k, "Value:", v)
	}

	println("length:", len(map1))

	// map2 := make(map[string]string) // shorthand declaration of map

	map2 := map1 // shallow copy

	map2["560086"] = "Bangalore-2"

	fmt.Println(map1)

	delete(map2, "560086")

	fmt.Println(map1)

	clear(map2)
	fmt.Println(map1)
	println("length:", len(map1))

}

// maps
// valus are mapped to keys
// O(1)
// maps in golang are not ordered
// hashfunc - > does mathematical func on the key

// 597a843f7613037df9a 7644459e295c188b870f3
// 3						4
// e08d99312033ce18299 e10c77822ec28a2bf7560
// 3						5
