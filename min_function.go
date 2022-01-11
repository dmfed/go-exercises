package main

import "fmt"

func min_rec(n ...int) int {
	if len(n) == 1 {
		return n[0]
	}
	a := n[0]
	b := min_rec(n[1:]...)
	if b < a {
		a = b
	}
	return a
}

func min_iter(n ...int) int {
	var m int
	if len(n) == 0 {
		return m
	}
	m = n[0]
	for x := range n {
		if n[x] < m {
			m = n[x]
		}
	}
	return m
}

func main() {
	n := []int{12, 13, 14, 15, 3, 4, 2, 1}
	fmt.Println(min_rec(n...))
	fmt.Println(min_iter(4, 2, 6, 1))
}
