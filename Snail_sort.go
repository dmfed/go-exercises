package main

/*
Snail Sort
Given an n x n array, return the array elements arranged from
outermost elements to the middle element, traveling clockwise.

array = [[1,2,3],
         [4,5,6],
         [7,8,9]]
snail(array) #=> [1,2,3,6,9,8,7,4,5]
For better understanding, please follow the numbers of the next
array consecutively:

array = [[1,2,3],
         [8,9,4],
         [7,6,5]]
snail(array) #=> [1,2,3,4,5,6,7,8,9]
This image will illustrate things more clearly:


NOTE: The idea is not sort the elements from the lowest value
to the highest; the idea is to traverse the 2-d array in a clockwise snailshell pattern.
*/

import "fmt"

func Snail(arr [][]int) []int {
	out := []int{}
	if len(arr[0]) == 0 {
		return out
	}
	offset := 0
	for offset <= len(arr)/2 {
		for i := offset; i < len(arr)-1-offset; i++ {
			out = append(out, arr[offset][i])
		}
		for i := offset; i < len(arr)-1-offset; i++ {
			out = append(out, arr[i][len(arr)-1-offset])
		}
		for i := offset; i < len(arr)-1-offset; i++ {
			out = append(out, arr[len(arr)-offset-1][len(arr)-i-1])
		}
		for i := offset; i < len(arr)-1-offset; i++ {
			out = append(out, arr[len(arr)-i-1][offset])
		}
		offset++
	}
	if len(arr)%2 > 0 {
		out = append(out, arr[offset-1][offset-1])
	}
	return out
}

func main() {
	test1 := [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}
	test2 := [][]int{{1, 2, 3, 4, 5},
		{6, 7, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	}
	fmt.Println(Snail(test1))
	fmt.Println(Snail(test2))
}
