package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
    "unicode"
)

func main() {
	phrase := readString()
    words := strings.Fields(phrase)
    var abbr string
    for _, w := range words{
        letter := []rune(w)
        if unicode.IsLetter(letter[0]){
            letter := unicode.ToUpper(letter[0])
            str := string(letter)
            abbr = abbr + str
        }
    }

	//    ее к верхнему регистру через `unicode.ToUpper()`
	// 3. Если слово начинается не с буквы, игнорируйте его
	//    проверяйте через `unicode.IsLetter()`
	// 4. Составьте слово из получившихся букв и запишите его
	//    в переменную `abbr`
    // ...

	fmt.Println(string(abbr))
}

func readString() string {
	rdr := bufio.NewReader(os.Stdin)
	str, _ := rdr.ReadString('\n')
	return str
}
