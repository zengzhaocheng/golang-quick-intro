package main

import (
	"fmt"
	"unsafe"
)

func main() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str" //多重赋值
	area = LENGTH * WIDTH
	fmt.Printf("area is : %d", area)
	println()
	println(a, b, c)

	const (
		Unknown = 0
		Female  = 1
		Male    = 2
	)

	const (
		e = "abc"
		f = len(e)
		g = unsafe.Sizeof(e)
	)
	println(e, f, g)
	const (
		a1 = iota
		a2 = iota
		a3 = iota
	)
	println(a1, a2, a3)
	const (
		b1 = iota
		b2
		b3
		c1 = "ha"
		c2
		c3 = 100
		c4
		c5 = iota
		c6
	)
	println(b1, b2, b3, c1, c2, c3, c4, c5, c6)

	const (
		i = 1 << iota
		j = 3 << iota
		k
		l
	)

	fmt.Println("i = ", i)
	fmt.Println("j = ", j)
	fmt.Println("k = ", k)
	fmt.Println("l = ", l)

}
