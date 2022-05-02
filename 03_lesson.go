package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	// foo1()
	foo2()
}

func foo1() {
	var s string
	fmt.Scan(&s)
	d, _ := time.ParseDuration(s)
	fmt.Println(s, "=", d.Minutes(), "min")
}

func foo2() {
	var x1 float64
	var y1 float64
	var x2 float64
	var y2 float64

	fmt.Scan(&x1, &y1, &x2, &y2)

	d := math.Sqrt(math.Pow((x1-x2), 2) + math.Pow((y1-y2), 2))

	fmt.Println(d)
}
