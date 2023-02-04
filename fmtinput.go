package main

import (
	"fmt"
)

func main() {
	var stockcode = 123
	var enddate = "2023-02-04"
	var url = "Code = %d & end Date = %s"
	fmt.Printf(url, stockcode, enddate)
}
