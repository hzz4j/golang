package main

import "fmt"
import "strconv"

func main() {
	fmt.Printf("Please input a float number:")
	var fNum float32
	fmt.Scan(&fNum)
	fmt.Printf("%d \n", int(fNum))
	s := strconv.FormatInt(-42, 10)
	fmt.Println(s)
}
