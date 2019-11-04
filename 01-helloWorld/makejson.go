package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("create a map and print json")

	fmt.Printf("please input user name: ")
	var name string
	fmt.Scan(&name)

	var address string
	fmt.Printf("please input user name: ")
	fmt.Scan(&address)
	person := map[string]string{
		"name":    name,
		"address": address}

	if res, err := json.Marshal(person); err == nil {
		fmt.Println(string(res))
	}
}

/**
create a map and print json
please input user name: Q10Viking
please input user name: BeiJing
{"address":"BeiJing","name":"Q10Viking"}
*/
