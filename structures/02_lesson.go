package main

import (
    "fmt"
)

type person struct {
    name string
    age  int
    sex any
}

type book struct {
    title  string
    author person
}


func (p person) sayHello()  {
    fmt.Println("Hello, my name is", p.name, " and I'm", p.age, "years old")
}


func main()  {

    anton := newPerson("Anton")
    fmt.Println(anton)

    anton.sayHello()

    mybook := newBook("IT", "King")
    fmt.Println(mybook)


}




func newPerson(name string) *person {
    p := person{name: name}
    p.age = 42
    p.sex = nil
    return &p




}


func newBook(title string, author string) *book {
    b := book{title: title}
    b.author = *newPerson(author)
    return &b
}
