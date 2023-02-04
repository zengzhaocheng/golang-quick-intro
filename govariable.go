package main

import "fmt"

func main() {
	var a string = "zen"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	b, c = c, b //牛逼！

	fmt.Println(b, c)
	//var a int 不能重复定义
	var d int
	fmt.Println(d)

	var e bool
	fmt.Println(e)

	var f float64
	var s string
	v := "xiexie"
	fmt.Printf("%v %v %q\n", f, v, s)

	//类型相同的多个变量
	/*
		var vname1, vname2, vname3 int
		vname1, vname2, vname3 = 1, 2, 3

		var n1, n2 = 1.2, 23.3 // 自动推断
		test1, test2 := "hehe", "haha"
	*/
}

//因式分解的写法 一般用于声明全局变量
var (
	var1 int
	var2 float32
)
