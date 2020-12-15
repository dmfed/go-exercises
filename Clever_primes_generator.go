package main

import (
	"fmt"
)

var iterations int = 0

type primes []int

func newPrimes() *primes {
	var p primes = []int{2}
	return &p
}

func (p *primes) next() int {
	result := (*p)[len(*p)-1]
	n := result + 1
	found := false
	for !found {
		found = true
		for x := 2; x * x <= n; x++ {
			iterations++
			if n%x == 0 {
				found = false
				n++
				break
			}
		}
		if found {
			*p = append(*p, n)
		}
	}
	return result
}

func main() {
	b := newPrimes()
	for i := 0; i < 10; i++ {
		n := b.next()
		fmt.Println(n)
	}
	fmt.Printf("Second found in %d interations", iterations)
}
