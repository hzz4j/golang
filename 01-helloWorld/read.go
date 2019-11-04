package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {

	type Person struct {
		Fname string `json:"fname"`
		Lname string `json:"lname"`
	}

	var filename string
	var result []Person
	fmt.Printf("please input filename: ")
	fmt.Scan(&filename)
	fhandle, _ := os.Open(filename)

	fileScanner := bufio.NewScanner(fhandle)
	for fileScanner.Scan() {
		name := strings.Split(fileScanner.Text(), " ")
		p := Person{name[0], name[1]}
		result = append(result, p)
	}
	fhandle.Close()
	for _, value := range result {
		fmt.Println(value)
		data, _ := json.Marshal(value)

		fmt.Println(string(data))
	}
}
