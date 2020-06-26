/*
Input

a string strng of n positive numbers (n = 0 or n >= 2)
Let us call weight of a number the sum of its digits. For example 99
will have "weight" 18, 100 will have "weight" 1.

Two numbers are "close" if the difference of their weights is small.

Task:
For each number in strng calculate its "weight" and then find two
numbers of strng that have:

the smallest difference of weights ie that are the closest
with the smallest weights
and with the smallest indices (or ranks, numbered from 0) in strng
Output:

an array of two arrays, each subarray in the following format:
[number-weight, index in strng of the corresponding number, original
corresponding number instrng]

or a pair of two subarrays (Haskell, Clojure, FSharp) or an array of
tuples (Elixir, C++)

or a (char*) in C mimicking an array of two subarrays or a string

or a matrix in R (2 rows, 3 columns, no columns names)

The two subarrays are sorted in ascending order by their number weights
if these weights are different, by their indexes in the string if they
have the same weights.

Examples:
Let us call that function closest

strng = "103 123 4444 99 2000"
the weights are 4, 6, 16, 18, 2 (ie 2, 4, 6, 16, 18)

closest should return [[2, 4, 2000], [4, 0, 103]] (or ([2, 4, 2000], [4, 0, 103])
or [{2, 4, 2000}, {4, 0, 103}] or ... depending on the language)
because 2000 and 103 have for weight 2 and 4, ther indexes in strng are 4 and 0.
The smallest difference is 2.
4 (for 103) and 6 (for 123) have a difference of 2 too but they are not
the smallest ones with a difference of 2 between their weights.
....................

strng = "80 71 62 53"
All the weights are 8.
closest should return [[8, 0, 80], [8, 1, 71]]
71 and 62 have also:
- the smallest weights (which is 8 for all)
- the smallest difference of weights (which is 0 for all pairs)
- but not the smallest indices in strng.
*/

package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func find_weights(s string) [][]int {
	nums := strings.Fields(s)
	weights := make([][]int, 0)
	for i, val := range nums {
		weight := 0
		for _, num := range val {
			toadd, _ := strconv.Atoi(string(num))
			weight += toadd
		}
		num, _ := strconv.Atoi(val)
		weights = append(weights, []int{weight, i, num})
	}
	sort.Slice(weights, func(i, j int) bool {
		if weights[i][0] == weights[j][0] {
			return weights[i][1] < weights[j][1]
		}
		return weights[i][0] < weights[j][0]
	})
	return weights
}

func Closest(strng string) string {
	weights := find_weights(strng)
	if len(weights) == 0 {
		return "[(), ()]"
	}
	min1, min2 := weights[0], weights[1]
	for i := 2; i < len(weights); i++ {
		if (min2[0] - min1[0]) > (weights[i][0] - weights[i-1][0]) {
			min1, min2 = weights[i-1], weights[i]
		}
	}
	return fmt.Sprintf("[(%v, %v, %v), (%v, %v, %v)]", min1[0], min1[1], min1[2], min2[0], min2[1], min2[2])
}

func main() {
	fmt.Println(Closest("239382 162 254765 182 485944 134 468751 62 49780 108 54"))
}
