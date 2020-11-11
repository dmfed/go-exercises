/*
In this Kata, you will be given an integer array and your task is to return the sum of elements occupying prime-numbered indices.

The first element of the array is at index 0.

Good luck!
*/

package main

type primes struct {
	have []int
}

func newPrimes() *primes {
	return &primes{have: []int{2}}
}

func (p *primes) next() int {
	res := p.have[len(p.have)-1]
	n := res + 1
	found := false
	for !found {
		found = true
		for _, num := range p.have {
			if n%num == 0 {
				found = false
				n++
				break
			}
		}
		if found {
			p.have = append(p.have, n)
		}
	}
	return res
}

func Solve(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	sum := 0
	pr := newPrimes()
	for {
		n := pr.next()
		if n < len(arr) {
			sum += arr[n]
		} else {
			break
		}
	}
	return sum
}
