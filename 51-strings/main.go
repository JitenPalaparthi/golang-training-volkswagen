package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	fruits := []string{"banana", "apple", "orange", "mango"}
	fmt.Println(fruits)
	sort.Strings(fruits)
	fmt.Println(fruits)

	str := "Hello World,How are you doing,whats up,what is happening over there\n"

	strs := strings.Split(str, ",")

	for _, str := range strs {
		println(str)
	}

}
