package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Printf("input string:   ")
	fmt.Scan(&input)
	if strings.HasPrefix(input, "i") && strings.HasSuffix(input, "n") && strings.Contains(input,"a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
