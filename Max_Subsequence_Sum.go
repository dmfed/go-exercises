/*
 The maximum sum subarray problem consists in finding the maximum sum
 of a contiguous subsequence in an array or list of integers:

maxSequence [-2, 1, -3, 4, -1, 2, 1, -5, 4]
-- should be 6: [4, -1, 2, 1]
Easy case is when the list is made up of only positive numbers and the
maximum sum is the sum of the whole array. If the list is made up of
only negative numbers, return 0 instead.

Empty list is considered to have zero greatest sum. Note that the empty
list or array is also a valid sublist/subarray.
*/

package main

import "fmt"

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaximumSubarraySum(numbers []int) int {
	max, curr := 0, 0
	for i := 0; i < len(numbers); i++ {
		if numbers[i] >= curr+numbers[i] {
			curr = numbers[i]
		} else {
			curr += numbers[i]
		}
		max = Max(max, curr)
	}
	return max
}

func main() {
	a := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(MaximumSubarraySum(a))
}

/* Красивое решение, но не мое :)
 func max(a, b int) int {
  if a > b {
    return a
  } else {
    return b
  }
}

func MaximumSubarraySum(numbers []int) int {
  res, sum := 0, 0
  for i := range numbers {
    sum += numbers[i]
    res = max(res, sum)
    sum = max(sum, 0)
  }
  return res
}
*/
