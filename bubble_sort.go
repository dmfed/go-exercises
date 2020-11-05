package main

import (
	"fmt"
	"math/rand"
	"time"
)

// BubbleSort1 is the most trivial implementation of Bubble Sort
func BubbleSort1(arr []int) {
	passes := 0
	swaps := 0
	totalsteps := 0
	for {
		passes++
		swapped := false
		for i := 1; i < len(arr); i++ {
			totalsteps++
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				swaps++
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	fmt.Printf("Made %d passes and made %d swaps\nTotal steps taken: %v\n", passes, swaps, totalsteps)
}

// BubbleSort2 differs in that after each pass in stops
// visiting the last cells of array (which are guaranteed
// to be sorted by the time)
func BubbleSort2(arr []int) {
	passes := 0
	totalsteps := 0
	swaps := 0
	n := len(arr)
	for n >= 1 {
		passes++
		swapped := false
		for i := 1; i < n; i++ {
			totalsteps++
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				swapped = true
				swaps++
			}
		}
		n--
		if !swapped {
			break
		}
	}
	fmt.Printf("Made %d passes and made %d swaps\nTotal steps taken: %v\n", passes, swaps, totalsteps)
}

// BubbleSort3 does not visit end of the array if it happens to be sorted already
func BubbleSort3(arr []int) {
	passes := 0
	swaps := 0
	totalsteps := 0
	n := len(arr)
	for n >= 1 {
		passes++
		newn := 0
		for i := 1; i < n; i++ {
			totalsteps++
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				newn = i
				swaps++
			}
		}
		n = newn
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
	// a := []int{11, 2, 1, 7, 3, 4, 6, 7, 8, 9, 10}
	a := generateRandomArray(10)

	b := copyArray(a)
	fmt.Println("Incoming array for BubbleSort1()", b)
	BubbleSort1(b)
	fmt.Println("Result of BubbleSort1()", b, "\n")

	c := copyArray(a)
	fmt.Println("Incoming array for BubbleSort2()", c)
	BubbleSort2(c)
	fmt.Println("Result of BubbleSort2()", c, "\n")

	d := copyArray(a)
	fmt.Println("Incoming array for BubbleSort3()", d)
	BubbleSort3(d)
	fmt.Println("Result of BubbleSort3()", d, "\n")
}
