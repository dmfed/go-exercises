package main

import (
	"fmt"
)

func permutations_k_length(list []int, k int) [][]int {
	if len(list) == 1 {
		out := make([][]int, 1)
		out[0] = list
		return out
	}
	out := make([][]int, 0)
	for i := 0; i < len(list); i++ {
		first := list[i]
		tmp := make([]int, len(list)-1)
		copy(tmp, list[:i]) // This is ugly. Needs a better approach...
		copy(tmp[i:], list[i+1:])
		for _, item := range permutations_k_length(tmp, k) {
			out = append(out, append([]int{first}, item...))
		}
	}
	for _, val := range out {
		val = val[:k]
	}
	return out
}

func main() {
	fmt.Println(permutations_k_length([]int{1, 2, 3}, 2))
}
