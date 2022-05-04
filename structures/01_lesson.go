package main

import (
    "fmt"
)

type person struct {
    name string
    age  int
    height float32
}


func main()  {

    anton := person{}
    fmt.Println("default values >>> ", anton)
    anton.name = "Anton"
    anton.age = 30
    anton.height = 1.75
    fmt.Println("after define params >>> ", anton)



}

func newPerson(name string) *person {
    p := person{name: name}
    p.age = 42
    return &p
}
