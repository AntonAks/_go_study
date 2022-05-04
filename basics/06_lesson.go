package main

import(
  "fmt"
)



func main()  {
  simple_if()
  cases()
  lang_case_function()
}


func simple_if()  {
  if 7%2 == 0 {
      fmt.Println("7 is even")
  } else {
      fmt.Println("7 is odd")
  }
  // 7 is odd


  if 8%4 == 0 {
      fmt.Println("8 is divisible by 4")
  }
  // 8 is divisible by 4


  if num := 9; num < 0 {
    fmt.Println(num, "is negative")
      } else if num < 10 {
    fmt.Println(num, "has 1 digit")
      } else {
    fmt.Println(num, "has multiple digits")
  }
}


func cases(){
  i := 2
  fmt.Print("Write ", i, " as ")
  switch i {
    case 1:
      fmt.Println("one")
    case 2:
      fmt.Println("two")
    case 3:
      fmt.Println("three")}
}


func lang_case_function(){
  var code string
	fmt.Scan(&code)
  switch code {
  case "en":
    fmt.Println("English")
  case "fr":
    fmt.Println("French")
  case "ru", "rus":
    fmt.Println("Russian")
  default: fmt.Println("Unknown")
  }
}
