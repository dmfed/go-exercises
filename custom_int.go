package main

import (
	"fmt"
)

type myint int

func (n myint) Report() string {
	return fmt.Sprintf("I'm number %d", n)
}

func main() {
	a := myint(6)
	fmt.Println(a)
	fmt.Println(a.Report(), 5+a)
}
