package main

import (
	"fmt"
	"sort"
	"strconv"
)

/**
输入数字到slice中，并排序
*/
func main() {
	var value string
	valSlice := make([]int, 0, 3)
	for {
		fmt.Println("please input a number,or input 'x' to exist:   ")
		fmt.Scan(&value)
		if i, err := strconv.Atoi(value); err == nil {
			valSlice = append(valSlice, i)
		}
		if "x" == value {
			fmt.Println("Bye!")
			break
		}
	}
	fmt.Printf("Before sort the result of slice is: \n %v %d \n", valSlice, len(valSlice))
	sort.Ints(valSlice)
	fmt.Printf("After sort the result of slice is: \n %v %d", valSlice, len(valSlice))
}
