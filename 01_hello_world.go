package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	basicTypes()
	compositeTypes()
}

func basicTypes() {

	b := false
	s := `hello`
	i := 1
	f64 := 3.14

	var r rune = 'R'

	fmt.Println(`basic >>> `, b, s, i, f64, r)
}

func compositeTypes() {

	//
	nums := [7]int{1, 3, 5, 7, 2, 3, 4}
	fmt.Println(`nums >>> `, nums)

	//
	words := []string{`hello`, `big`, `world`}
	fmt.Println(`words >>> `, words)

	//
	words = append(words, `GO`, `LANG`)
	fmt.Println(`words (append)>>> `, words)

	//
	dict := map[string]int{"one": 1, "two": 2}
	fmt.Println(dict)
	fmt.Println(dict["one"])

	//
	type person struct {
		name string
		age  int
	}

	p := person{name: "Mike", age: 34}
	fmt.Printf("Name: %s, age: %d\n", p.name, p.age)

}
