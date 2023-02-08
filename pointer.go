package main

import "fmt"

const Max int = 3

func main() {
	var a int = 20
	var ip *int
	ip = &a
	fmt.Printf("a 变量地址：%x\n", &a)
	fmt.Printf("ip 存储的指针地址： %x \n", ip)
	fmt.Printf("*ip 变量的值 %d\n", *ip)

	var ptr *int
	fmt.Printf("ptr的值为 ： %x\n", ptr)
	if ptr != nil {
		fmt.Printf("ptr不是空指针")
	} else {
		fmt.Printf("ptr是空指针")
	}
	println("======指针数组=======")
	b := []int{10, 100, 1000}
	var i int
	var ptr1 [Max]*int

	for i = 0; i < Max; i++ {
		fmt.Printf("b[%d] = %d\n", i, b[i])
	}

	println("整数地址赋值给指针数组")
	for i = 0; i < Max; i++ {
		ptr1[i] = &b[i]
	}
	for i = 0; i < Max; i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptr1[i])
	}

	println("========指向指针的指针========")
	var aa int
	var ptr2 *int
	var pptr **int
	aa = 3000
	ptr2 = &aa
	pptr = &ptr2
	fmt.Printf("变量 aa = %d\n", aa)
	fmt.Printf("指针变量 *ptr = %d\n", *ptr2)
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)

}
