/*
In genetic algorithms, a crossover allows 2 chromosomes to exchange part of their genes.

For more details, you can visit these wikipedia links : Genetic algorithm and Crossover.

A chromosome is represented by a list of genes.
Consider for instance 2 chromosomes xs (with genes [x,x,x,x,x,x]) and ys (with genes [y,y,y,y,y,y])

A single-point crossover at index 2 would give :

new chromosome1 = []int { x,x,y,y,y,y } and new chromosome2 = []int { y,y,x,x,x,x }
A two-point crossover at indices 1 and 3 would give :

new chromosome1 = []int { x,y,y,x,x,x } and new chromosome2 = []int { y,x,x,y,y,y }
We want to generalize this to the notion of n-point crossover, and create a generic function with
the following signature :

func Crossover(ns []int, xs []int,ys []int) ([]int, []int)
where ns would be a list of cross-point indices.

We could compute a triple-point crossover of 2 chromosomes xs and ys the following way :

Crossover(int[] { 2, 5, 21 }, xs, ys)
The transformed first chromosome must appear at the first position of the tuple. the second one
at the second position. Therefore :

Crossover([]int { 1 }, []int {1,2,3}, []int {10,11,12}) should be ([]int { 1,11,12 }, int[] { 10,2,3 })
If a cross-point index is repeated, it should be considered only once. Indices can also be
provided unordered, so your function could be called with the following indices :

Crossover(int[] { 7,5,3,5 }, xs, ys)
In this case, you would have to consider only indices [7,5,3] and deal with the fact they are unordered.

Chromosome indices are 0-based and you can also assume that :

(length xs) == (length ys)
crossover indices will never exceed the length of xs or ys.
*/

package main

import (
	"sort"
)

// ns : slice of indices
// xs, ys : chromosomes as slices of ints
func Crossover(ns []int, xs []int, ys []int) ([]int, []int) {
	sort.Ints(ns)
	index := 0
	for index < len(ns) {
		currindex := ns[index]
		for i := ns[index]; i < len(xs); i++ {
			xs[i], ys[i] = ys[i], xs[i]
		}
		for index < len(ns) && currindex == ns[index] {
			index++
		}
	}
	return xs, ys
}
