package main

import (
	"fmt"
)

func fibsUpTo(limit int) chan int {
      out := make(chan int)
      go func() {
			a, b := 1, 2
			for a <= limit {
				out <- a
				a, b = b, a+b                
            }
            close(out)
      }()
      return out
}

func main() {
	for v := range fibsUpTo(10) {
		fmt.Println(v)
	}
}
