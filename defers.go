// Testing defers in Go
package main

import (
	"fmt"
)

func worker(a int, b int) int {
	c := 0
	c = a + b
	defer fmt.Println("Deferred:", c)
	fmt.Println("Inner:", c)
	c--
	fmt.Println("Inner2:", c)
	return c
}

func main() {
	fmt.Println("Returned", worker(3, 4))

}
