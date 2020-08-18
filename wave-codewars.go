/*
In this simple Kata your task is to create a function that turns a string into a
Mexican Wave. You will be passed a string and you must return that string in an
array where an uppercase letter is a person standing up.
Rules

1.  The input string will always be lower case but maybe empty.
2.  If the character in the string is whitespace then pass over it as if it
was an empty seat.

Example
wave("hello") => []string{"Hello", "hEllo", "heLlo", "helLo
*/

package main

import (
	"fmt"
	"unicode"
)

func wave(words string) []string {
	out := []string{}
	tmp := []rune(words)
	for i := 0; i < len(tmp); i++ {
		if unicode.IsSpace(tmp[i]) {
			continue
		}
		out = append(out, words[:i]+string(unicode.ToUpper(tmp[i]))+words[i+1:])
	}
	return out
}

func main() {
	fmt.Println(wave("hello"))
	fmt.Println(wave("he llo"))
}
