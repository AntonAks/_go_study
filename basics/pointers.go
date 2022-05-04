package main

import (
    "fmt"
)


func main()  {
    var iptr *int
    i := 42
    iptr = &i

    fmt.Println(iptr)
}
