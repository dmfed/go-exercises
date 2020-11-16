package main

import (
        "fmt"
        )

type primes []int

func newPrimes() *primes {
  var p primes = []int{2}
  return &p
}

func (p *primes) next() int {
    res := (*p)[len(*p)-1]
    n := res + 1
    found := false
    for !found {
      found = true
      for _, num := range *p {
        if n % num == 0 {
          found = false
          n++
          break
        }
      }
      if found {
        *p = append(*p, n)
      }
    }
    return res
}

func main() {
  b := newPrimes()
	for i := 0; i < 10; i++ {
    n := b.next()
    fmt.Println(n)
	}
}
