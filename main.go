package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("app.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(data))
	for a, b := range strings.Split(string(data), " ") {
		fmt.Printf("i: %d\nword: %s\n", a, string(b))
	}
}
