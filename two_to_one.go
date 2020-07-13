/*
Take 2 strings s1 and s2 including only letters from ato z.
Return a new sorted string, the longest possible, containing distinct letters,

each taken only once - coming from s1 or s2.
Examples:
a = "xyaabbbccccdefww"
b = "xxxxyyyyabklmopq"
longest(a, b) -> "abcdefklmopqwxy"

a = "abcdefghijklmnopqrstuvwxyz"
longest(a, a) -> "abcdefghijklmnopqrstuvwxyz"
*/

package main

import (
	"fmt"
	"sort"
	"strings"
	)

func TwoToOne(s1 string, s2 string) string {
	set := make(map[string]bool)
	for _, letter := range s1 {
		set[string(letter)] = true
	}
	for _, letter := range s2 {
		set[string(letter)] = true
	}
	result := []string{}
	for letter := range set {
		result = append(result, letter)
	}
	sort.Slice(result, func(i, j int) bool {return result[i]<result[j]})
	return strings.Join(result, "")
}

func main() {
	fmt.Println(TwoToOne("loopingisfunbutdangerous", "lessdangerousthancoding") == "abcdefghilnoprstu")
}
