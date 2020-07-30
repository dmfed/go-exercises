package main

import "fmt"

func rwChan(ch chan string) {
	for {
		a := <-ch
		ch <- a
	}
}

func main() {
	twoway := make(chan string)
	go rwChan(twoway)
	count := 0
	for count < 10 {
		twoway <- "hello"
		fmt.Println(<-twoway)
		count++
	}
}
