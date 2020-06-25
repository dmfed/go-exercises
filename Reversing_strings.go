package main

import (
	"fmt"
	"strings"
)

const (
	s string = "abcd\nefgh\nijkl\nmnop"
	// Below are for Vert
	sv1 string = "hSgdHQ\nHnDMao\nClNNxX\niRvxxH\nbqTVvA\nwvSyRu"
	sv2 string = "QHdgSh\noaMDnH\nXxNNlC\nHxxvRi\nAvVTqb\nuRySvw"
	// These are for Horiz
	sh1 string = "lVHt\nJVhv\nCSbg\nyeCt"
	sh2 string = "yeCt\nCSbg\nJVhv\nlVHt"
)

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func VertMirror(s string) string {
	groups := strings.Split(s, "\n")
	for i := 0; i < len(groups); i++ {
		groups[i] = Reverse(groups[i])
	}
	return strings.Join(groups, "\n")
}
func HorMirror(s string) string {
	groups := strings.Split(s, "\n")
	for i, j := 0, len(groups)-1; i < len(groups)/2; i, j = i+1, j-1 {
		groups[i], groups[j] = groups[j], groups[i]
	}
	return strings.Join(groups, "\n")
}

type FParam func(string) string

func Oper(f FParam, x string) string {
	a := f(x)
	return a
}

func main() {
	fmt.Println(s, "\n\n", Oper(VertMirror, s), "\n\n", Oper(HorMirror, s), "\n\n")
	fmt.Println(sh1, "\n\n")
	fmt.Println(HorMirror(sh1), "\n\n")
	fmt.Println(sv1, "\n\n")
	fmt.Println(Oper(VertMirror, sv1), "\n\n")
}
