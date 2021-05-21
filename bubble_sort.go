package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// BubbleSort1 is the most trivial implementation of Bubble Sort
func BubbleSort1(arr []int) {
	passes := 0
	swaps := 0
	totalsteps := 0
	swapped := true
	for swapped {
		passes++
		swapped = false
		for i := 1; i < len(arr); i++ {
			totalsteps++
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				swaps++
				swapped = true
			}
		}
	}
	fmt.Printf("Made %d passes and made %d swaps\nTotal steps taken: %v\n", passes, swaps, totalsteps)
}

// BubbleSort2 differs in that after each pass it stops
// visiting the last cells of array (which are guaranteed
// to be sorted by the time)
func BubbleSort2(arr []int) {
	passes := 0
	totalsteps := 0
	swaps := 0
	end := len(arr)
	for end > 0 {
		passes++
		for i := 1; i < end; i++ {
			totalsteps++
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				swaps++
			}
		}
		end--
	}
	fmt.Printf("Made %d passes and made %d swaps\nTotal steps taken: %v\n", passes, swaps, totalsteps)
}

// BubbleSort2pluscombines techniques from BubbleSort() and BubbleSort2()
func BubbleSort2plus(arr []int) {
	passes := 0
	totalsteps := 0
	swaps := 0
	end := len(arr)
	swapped := true
	for swapped {
		swapped = false
		passes++
		for i := 1; i < end; i++ {
			totalsteps++
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				swapped = true
				swaps++
			}
		}
		end--
	}
	fmt.Printf("Made %d passes and made %d swaps\nTotal steps taken: %v\n", passes, swaps, totalsteps)
}

// BubbleSort3 does not visit end of the array if it happens to be sorted already
func BubbleSort3(arr []int) {
	passes := 0
	swaps := 0
	totalsteps := 0
	end := len(arr)
	for end > 0 {
		passes++
		newend := 0
		for i := 1; i < end; i++ {
			totalsteps++
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				newend = i
				swaps++
			}
		}
		end = newend
	}
	fmt.Printf("Made %d passes and made %d swaps\nTotal steps taken: %v\n", passes, swaps, totalsteps)
}

func generateRandomArray(length int) []int {
	arr := []int{}
	for i := 0; i < length; i++ {
		arr = append(arr, rand.Intn(421))
	}
	return arr
}

func copyArray(in []int) []int {
	out := make([]int, len(in))
	copy(out, in)
	return out
}

func main() {
	rand.Seed(time.Now().UnixNano())
	//a := []int{11, 2, 1, 7, 3, 4, 6, 5, 8, 9, 10} // Example demonstrates efficiency of third impl.
	a := generateRandomArray(10)

	b := copyArray(a)
	fmt.Println("Incoming array for BubbleSort1()", b)
	BubbleSort1(b)
	fmt.Println("Result of BubbleSort1()", b, "\n")

	c := copyArray(a)
	fmt.Println("Incoming array for BubbleSort2()", c)
	BubbleSort2(c)
	fmt.Println("Result of BubbleSort2()", c, "\n")

	f := copyArray(a)
	fmt.Println("Incoming array for BubbleSort2plus()", c)
	BubbleSort2plus(f)
	fmt.Println("Result of BubbleSort2plus()", f, "\n")

	d := copyArray(a)
	fmt.Println("Incoming array for BubbleSort3()", d)
	BubbleSort3(d)
	fmt.Println("Result of BubbleSort3()", d, "\n")

	e := copyArray(a)
	fmt.Println("Incoming array for sort.Ints()", e)
	sort.Ints(e)
	fmt.Println("Result of sort.Ints()", e, "\n")

	for i := 0; i < len(b); i++ {
		if b[i] != e[i] || c[i] != e[i] || d[i] != e[i] || f[i] != e[i] {
			fmt.Println("Comparison with standard library failed")
		}
	}
}
