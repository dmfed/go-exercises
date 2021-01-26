package main

import (
	"fmt"
	"flag"
)

func FizzBuzz(n int) {
	for x := 1; x <= n; x++ {
		if x % 3 == 0 && x % 5 == 0 {
			fmt.Println("FizzBuzz")
		} else if x % 3 == 0 {
			fmt.Println("Fizz")
		} else if x % 5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(x)
		}
	}
}

func main() {
	n := flag.Int("n", 0, "Play FizzBuzz game up to n")
	flag.Parse()
	FizzBuzz(*n)
}
