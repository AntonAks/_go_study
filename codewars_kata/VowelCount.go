package main
import "fmt"
/*
Return the number (count) of vowels in the given string.
We will consider a, e, i, o, u as vowels for this Kata (but not y).
The input string will only consist of lower case letters and/or spaces.
*/

func main() {
    fmt.Println(GetCount("Hello World"))

}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func GetCount(str string) (count int) {
    vowels := []string{"a", "e", "i", "o", "u"}
    chars := []rune(str)
    count = 0
    for i:=0; i < len(chars); i++ {
        if contains(vowels, string(chars[i])){
            count++
            }
    }

    return count
}


func GetCount2(str string) (count int) {
  for _, c := range str {
    switch c {
    case 'a', 'e', 'i', 'o', 'u':
      count++
    }
  }
  return count
}
