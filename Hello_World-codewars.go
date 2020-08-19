package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func greet() string {
	URL := "https://raw.githubusercontent.com/dmfed/go-exercises/master/hello.txt"
	resp, err := http.Get(URL)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("Error", err)
		return ""
	}
	out, _ := ioutil.ReadAll(resp.Body)
	return string(out)
}

func main() {
	fmt.Println(greet())
}
