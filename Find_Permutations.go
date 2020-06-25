package main

import (
	"fmt"
)

func permutations(list []int) [][]int {
	if len(list) == 1 {
		out := make([][]int, 1)
		out[0] = list
		fmt.Printf("Returning %v\n", out)
		return out
	}
	out := make([][]int, 0)
	for i := 0; i < len(list); i++ {
		first := list[i]
		tmp := make([]int, len(list)-1)
		copy(tmp, list[:i])
		copy(tmp[i:], list[i+1:])
		for _, item := range permutations(tmp) {
			out = append(out, append([]int{first}, item...))
		}
	}
	return out
}

func main() {
	fmt.Println(permutations([]int{1, 2, 3}))
}
