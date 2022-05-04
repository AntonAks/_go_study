package main

import (
    "fmt"
)



func main()  {
    res := sumAbc(12, 15, 17)
    fmt.Println(res)

    r1, r2 := divide(130, 30)

    fmt.Println("r1:", r1, "r2:", r2)
}


func sumAbc(a, b, c int) int {
    return a + b + c
}

func divide(divisible, divisor int) (int, int) {
    quotient := divisible / divisor
    remainder := divisible % divisor
    return quotient, remainder
}
