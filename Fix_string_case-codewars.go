/*
n this Kata, you will be given a string that may have mixed uppercase and lowercase letters and your task is to convert that string to either lowercase only or uppercase only based on:

make as few changes as possible.
if the string contains equal number of uppercase and lowercase letters, convert the string to lowercase.
For example:

solve("coDe") = "code". Lowercase characters > uppercase. Change only the "D" to lowercase.
solve("CODe") = "CODE". Uppercase characters > lowecase. Change only the "e" to uppercase.
solve("coDE") = "code". Upper == lowercase. Change all to lowercase.
More examples in test cases. Good luck!
*/

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func solve(str string) string {
	tmp := []rune(str)
	lower, upper := 0, 0
	for _, r := range tmp {
		if unicode.IsLower(r) {
			lower++
		} else if unicode.IsUpper(r) {
			upper++
		}
	}
	if upper > lower {
		str = strings.ToUpper(str)
	} else {
		str = strings.ToLower(str)
	}
	return str
}

func main() {
	fmt.Println(solve("heLlo"))
	fmt.Println(solve("HELlo"))
	fmt.Println("GOgo")
}
