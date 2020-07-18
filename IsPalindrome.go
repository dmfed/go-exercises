/*
Write function isPalindrome that checks if a given string (case
insensitive) is a palindrome.

In Racket, the function is called palindrome?

(palindrome? "nope") ; returns #f
(palindrome? "Yay")  ; returns #t
*/
package main

import (
	"fmt"
	"strings"
)

func IsPalindrome(str string) bool {
	str = strings.ToLower(str)
	for i, j := 0, len(str)-1; i < len(str)/2; i, j = i+1, j-1 {
		if !(str[i] == str[j]) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsPalindrome("yayayay") == true)
	fmt.Println(IsPalindrome("hello") == false)
}
