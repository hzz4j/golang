1. 声明变量的类型不同，如float32,float64.
2. int32,int16类型的不同的变量，在赋值时需要进行转化,int32此时位方法名。
3. 两个重要的包`strconv`,`strings`
4. fmt.Scan需要一个地址作为参数
5. 4. Atoi(s) convert string to int

```go
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

```


```go
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

```

