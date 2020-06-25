package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, pos2 := adder(), adder()
	a := pos(1)
	b := pos2(2)
	c := pos(5)
	fmt.Println(a, b, c, pos(10))
}
