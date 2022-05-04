package main

import (
    "fmt"
)


func main()  {

    var iter = []int{1,2,3,4,5,6,7}
    fmt.Println(filter(predicate, iter))
}


func filter(predicate func(int) bool, iterable []int) []int {
    var _res = []int{}
    for _, i := range iterable{
        if predicate(i) {
            _res = append(_res, i)
        }
}
return _res
}

func predicate(i int) bool{
    fmt.Println(i, i%2==0)
    if i%2 == 0 {
        return true
    }
    return false
}
