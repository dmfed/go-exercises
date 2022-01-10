package main

import "fmt"

func min(n ...int) int {
	if len(n) == 1 {
		return n[0]
	}
	a := n[0]
	b := min(n[1:]...)
	if b < a {
		a = b
	}
	return a
}

func main() {
	n := []int{12, 13, 14, 15, 3, 4, 2, 1}
	fmt.Println(min(n...))
	fmt.Println(min(4, 2, 6, 1))
}
