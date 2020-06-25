package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func LCS(s1, s2 string) string {
	solution := make([][]int, len(s1)+1)
	for i, _ := range solution {
		solution[i] = make([]int, len(s2)+1)
	}
	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if s1[i-1] == s2[j-1] {
				solution[i][j] = solution[i-1][j-1] + 1
			} else {
				solution[i][j] = max(solution[i-1][j], solution[i][j-1])
			}
		}
	}
	answer := make([]rune, 0)
	for i, j := len(solution)-1, len(solution[0])-1; i > 0 && j > 0; {
		if solution[i-1][j-1] == solution[i][j]-1 && solution[i][j-1] == solution[i-1][j] && solution[i-1][j] == solution[i-1][j-1] {
			answer = append(answer, rune(s2[j-1]))
			i--
			j--
		} else if solution[i-1][j] >= solution[i][j-1] {
			i--
		} else if solution[i][j-1] > solution[i-1][j] {
			j--
		}
	}
	for i, j := 0, len(answer)-1; i < len(answer)/2; i, j = i+1, j-1 {
		answer[i], answer[j] = answer[j], answer[i]
	}
	return string(answer)
}

func main() {
	fmt.Println(LCS("132535365", "123456789"))
	fmt.Println(LCS("abcdef", "abc"))
}
