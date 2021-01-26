// tmp.go
//
// Copyright 2020 Dmitry Fedotov <dafedotov@gmail.com>
 

package main

import (
	"fmt"
	)
	
const (
    ColorBlack  	  = "\u001b[30m"
    ColorRed          = "\u001b[31m"
    ColorGreen        = "\u001b[32m"
    ColorYellow       = "\u001b[33m"
    ColorBlue         = "\u001b[34m"
    ColorReset        = "\u001b[0m"
)

func colorize(color string, message string) {
    fmt.Println(color, message, ColorReset)
}	
	
func main() {
	colorize(ColorRed, "hello")
}

