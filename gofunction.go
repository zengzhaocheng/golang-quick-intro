package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func getSequence() func() int {
	i := 1
	return func() int {
		i += 1
		return i
	}
}

func max(num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}
func swap(x, y string) (string, string) {
	return y, x
}

// 引用传递
func swap1(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

func app() func(string) string {
	t := "Hi"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	println(&t)
	return c
}

func main() {
	var a int = 100
	var b int = 200
	var ret int
	ret = max(a, b)
	fmt.Printf("最大值 ： %d\n", ret)
	c, d := swap("Google", "Runoob")
	fmt.Println(c, d)
	fmt.Printf("交换前，a = %d，b = %d\n", a, b)
	swap1(&a, &b)
	fmt.Printf("交换前，a = %d，b = %d\n", a, b)

	println("=====函数作为参数=====")
	//声明函数变量
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	//使用函数
	fmt.Println(getSquareRoot(9))

	println("====函数闭包（函数作为返回值）======")
	nextNumber := getSequence()

	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())

	println("===面试题1：https://polarisxu.studygolang.com/posts/go/action/go-closure/==== ")
	testa := app()
	testb := app()
	testa("go")
	fmt.Println(testb("All"))
	fmt.Println(testa("All"))

	println("=====方法=====")
	var c1 Circle
	c1.radius = 10.00
	fmt.Println("圆的面积 = ", c1.getArea())

}

// 方法就是包含接收者的函数
func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}
