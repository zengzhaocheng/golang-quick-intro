package main

import (
	"fmt"
	"reflect"
)

// 结构体
type Phone interface {
	call()
}
type NokiaPhone struct {
}

// 方法
func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func main() {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()

	println("======面试题1：https://juejin.cn/post/6960888943273312263=====")

	var v interface{}
	v = (*int)(nil)
	fmt.Println(v == nil)

	println("===========")
	var data *byte
	var in interface{}

	fmt.Println(data, data == nil)
	fmt.Println(in, in == nil)
	in = data
	fmt.Println(in, in == nil)

	var data1 *byte
	var in1 interface{}

	in1 = data1
	fmt.Println(IsNil(in1))

}

// 用了反射，什么是反射？
func IsNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}
