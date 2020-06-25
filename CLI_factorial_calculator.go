package main

import (
	"flag"
	"fmt"
)

func factorial(n int64) int64 {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	var n int64
	flag.Int64Var(&n, "n", -1, "I'll calculate factorial of this number")
	flag.Parse()
	if n < 0 {
		fmt.Println("We can only calculate factorial of n >= 0")
		flag.PrintDefaults()
	} else {
		fmt.Println(factorial(n))
	}
}
