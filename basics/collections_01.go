package main

import (
    "fmt"
    // "bytes"
    "strings"
    "strconv"
    
)


func main(){
    // var a [5]int
    // fmt.Println("empty:", a)
    // // empty: [0 0 0 0 0]
    //
    // a[4] = 100
    // fmt.Println("set:", a)
    // // set: [0 0 0 0 100]
    //
    // fmt.Println("get:", a[4], "len:", len(a))
    // // get: 100
    //
    // b := []int{1, 2, 3, 4, 5}
    // fmt.Println("init:", b)
    // // init: [1 2 3 4 5]
    // b = append(b, 3,4,5,6,7)
    // fmt.Println("append:", b)

    // exercise_1()
    // exercise_2()

}

func exercise_1()  {
    var text string
    var width int
    fmt.Scanf("%s %d", &text, &width)

    if len(text) >= width {
        text = text[:width] + "..."
    }

    fmt.Println(text)

}

func exercise_2()  {
    var number int
	fmt.Scanf("%d", &number)

	// Посчитайте, сколько раз каждая цифра встречается
	// в числе `number`. Запишите результат в карту `counter`,
	// где ключом является цифра числа, а значением -
	// сколько раз она встречается

	counter := make(map[int]int)
    number_s := strconv.Itoa(number)

    i := 1
    for i <= len(number_s) {
        a, _ := strconv.Atoi(string(number_s[i-1]))
        counter[a] = strings.Count(number_s, string(number_s[i-1]))
        i = i + 1
    }


    fmt.Println(counter)

}

func exercise_3()  {

}
