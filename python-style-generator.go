package main

import ("fmt"
	)

type fibchan chan int

func (f fibchan) Next() int {
	val, ok := <-f
	if ok {
		return val
	}
	return -1
}

func fibonacchi(limit int) fibchan {
	a, b := 0, 1
	out := make(chan int) 
	go func() {
		for i :=0; i < limit; i++ {
			out <- a
			a, b = b, a+b
		}
		close(out)
	}()
	return out
}

func main() {
	fib := fibonacchi(20)
	for i := 0; i < 5; i++ {
		fmt.Println("Next:", fib.Next())
	}
	for val := range fib {
		fmt.Println("Iterating:", val)
	}
}
