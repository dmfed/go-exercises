/*
The input is a string str of digits. Cut the string into chunks (a chunk here is a substring of the
initial string) of size sz (ignore the last chunk if its size is less than sz).

If a chunk represents an integer such as the sum of the cubes of its digits is divisible by 2, reverse
that chunk; otherwise rotate it to the left by one position. Put together these modified chunks and
return the result as a string.

If

sz is <= 0 or if str is empty return ""
sz is greater (>) than the length of str it is impossible to take a chunk of size sz hence return "".
Examples:
revrot("123456987654", 6) --> "234561876549"
revrot("123456987653", 6) --> "234561356789"
revrot("66443875", 4) --> "44668753"
revrot("66443875", 8) --> "64438756"
revrot("664438769", 8) --> "67834466"
revrot("123456779", 8) --> "23456771"
revrot("", 8) --> ""
revrot("123456779", 0) --> ""
revrot("563000655734469485", 4) --> "0365065073456944"
*/

package main

import (
	"fmt"
	"strings"
	"strconv"
	)

func Revrot(s string, n int) string {
	if len(s) == 0 || n > len(s) || n <= 0 {
		return ""
	}
	ints := []int{}
	for _, val := range s {
		n, _ := strconv.Atoi(string(val))
		ints = append(ints, n)
	}
	tmp := []string{}
	for i, cycles := 0, len(ints) / n; cycles > 0; i, cycles = i+n, cycles-1 {
		if sumCubes(ints[i:i+n]) % 2 == 0 {
			rev := reverseSlice(ints[i:i+n])
			for _, val := range rev {
				tmp = append(tmp, strconv.Itoa(val))
			}
		} else {
			rot := rotateSlice(ints[i:i+n])
			for _, val := range rot {
				tmp = append(tmp,  strconv.Itoa(val))
			}
		}
	}
	return strings.Join(tmp, "")
}

func sumCubes(s []int) int {
	out := 0
	for _, number := range s {
		out += number * number * number
	}
	return out
}

func reverseSlice(s []int) []int {
	for i, j := 0, len(s)-1; i < len(s)/2; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i] 
	}
	return s
}

func rotateSlice(s []int) []int {
	tmp := s[0]
	for i := 0; i < len(s)-1; i++ {
		s[i] = s[i+1]
	}
	s[len(s)-1] = tmp
	return s
}


func main() {
	fmt.Println(Revrot("733049910872815764", 5) == "330479108928157")
	fmt.Println(Revrot("123456987654", 6) == "234561876549")
	fmt.Println(Revrot("123456987653", 6) == "234561356789")
	fmt.Println(Revrot("66443875", 4) == "44668753")
}
