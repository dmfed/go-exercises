package main

import "fmt"

func main() {
	c := make(chan string, 2)
	c <- "hello"
	c <- "world"
	close(c)           // channel is closed here
	for s := range c { // we can still range over values already written
		fmt.Println(s)
	}
}
