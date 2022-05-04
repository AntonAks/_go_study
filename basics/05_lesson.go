package main

import (
	"fmt"
)

func main() {
	var source string
	var times int
	fmt.Scan(&source, &times)

	var result string = ""
	i := 1
	for i <= times {
		result = result + source
		i = i + 1
	}

	fmt.Println(result)
}
