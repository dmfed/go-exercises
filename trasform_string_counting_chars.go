package main

import (
        "fmt"
        "strconv"
        )

func transform(inp string) string {
	out := ""
	var prev rune
	count := 0
	for _, r := range inp {
		if r != prev {
			if count > 0 {
				out += strconv.Itoa(count)
			}
			out += string(r)
			prev = r
			count = 1
			continue
		}
		count++
	}
	if count > 0 {
		out += strconv.Itoa(count)
	}
	return out
}

func main() {
	a := "bbbbbbcccaabbbb"
	fmt.Println(transform(a))
}
