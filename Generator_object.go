// Generator_object.go
//
// Copyright 2020 Dmitry Fedotov <dafedotov@gmail.com>

package main

import (
	"fmt"
)

type InfiniteGenerator struct {
	arr     []string
	current int
}

func NewGenerator(slice []string) *InfiniteGenerator {
	g := InfiniteGenerator{arr: slice, current: 0}
	return &g
}
func (g *InfiniteGenerator) Next() string {
	val := g.arr[g.current]
	if g.current == len(g.arr)-1 {
		g.current = 0
	} else {
		g.current++
	}
	return val
}

func main() {
	g := NewGenerator([]string{"a", "b", "c"})
	for i := 0; i < 23; i++ {
		fmt.Println(g.Next())
	}
}
