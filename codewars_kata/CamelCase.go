package main
import (
    "fmt"
    ."strings"
)


/*
Write simple .camelCase method (camel_case function in PHP, CamelCase in C# or camelCase in Java) for strings.
All words must have their first letter capitalized without spaces.

For instance:
CamelCase("hello case")      // => "HelloCase"
CamelCase("camel case word") // => "CamelCaseWord"

Don't forget to rate this kata! Thanks :)
*/
func main(){
    fmt.Println(CamelCase("hello beautiful world"))
    fmt.Println(CamelCase2("hello beautiful world"))

}


func CamelCase(s string) string {

    words := Fields(s)
    result := ""
    for _, w := range words{
        result = result + Title(w)
    }

    return result
}



func CamelCase2(str string) string {
      return Join(Fields(Title(str)), "")
}
