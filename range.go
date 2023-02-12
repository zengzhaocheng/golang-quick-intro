package main

import (
	"fmt"
	"unsafe"
)

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	map1 := make(map[int]float32)
	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0
	map1[4] = 4.0

	// 读取 key 和 value
	for key, value := range map1 {
		fmt.Printf("key is: %d - value is: %f\n", key, value)
	}

	// 读取 key
	for key := range map1 {
		fmt.Printf("key is: %d\n", key)
	}

	// 读取 value
	for _, value := range map1 {
		fmt.Printf("value is: %f\n", value)
	}

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	for i, c := range "go" {
		fmt.Println(i, c)
	}

	println("========面试题1：https://kuricat.com/gist/golang-for-range-y6mtl======")
	i := 12
	arr := []int{1, 2, 3, 4}
	fmt.Println(unsafe.Pointer(&arr))
	for key := range arr {
		fmt.Println(unsafe.Pointer(&arr))
		arr = append(arr, 5)
		fmt.Printf("%+v", arr)
		if key > i { //并不会改变循环次数
			break
		}
	}
	println("========面试题2：https://segmentfault.com/a/1190000041327456======")
	var x = []string{"A", "B", "C"}
	fmt.Printf("address of x is : %p\n", x)
	for i, s := range x {
		print(i, s, "---")
		fmt.Printf("address of x is : %p\n", x)
		x[i+1] = "M"
		fmt.Printf("address of x is : %p\n", x)
		print(i, s, "---")
		x = append(x, "Z")
		fmt.Printf("address of x is : %p\n", x)
		x[i+1] = "Z"
		fmt.Printf("address of x is : %p\n", x)
		print(i, s, "---")
	}
	fmt.Printf("address of x is : %p\n", x)
	//slice做range遍历，Go编译器背后会做哪些事情？
	//slice什么时候扩容，扩容后的行为是怎么样的？
}
