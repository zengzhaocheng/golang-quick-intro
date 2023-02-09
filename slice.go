package main

import "fmt"

func main() {
	var numbers = make([]int, 3, 5) // len cap
	printSlice(numbers)

	println("===========空切片=============")
	var numbers1 []int
	printSlice(numbers1)
	if numbers1 == nil {
		fmt.Printf("empty slice")
	}
	println("===========切片截取=============")
	numbers2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers2)

	fmt.Println("numbers2 == ", numbers2)
	fmt.Println("numbers2[1:4]", numbers2[1:4])
	fmt.Println("numbers2[:3]", numbers2[:3])
	fmt.Println("numbers2[4:]", numbers2[4:])

	println("===========append() 和 copy() 函数=============")
	numbers = append(numbers, 0)
	printSlice(numbers)
	numbers = append(numbers, 1)
	printSlice(numbers)
	numbers = append(numbers, 3, 4, 5)
	printSlice(numbers)
	numbers3 := make([]int, len(numbers), (cap(numbers))*2)
	copy(numbers3, numbers)
	printSlice(numbers3)

	println("=====面试题1：https://juejin.cn/post/6844903859257606158====")
	s := []int{1, 2, 3, 4, 5, 6}
	//fmt.Printf("address of s:%v\n", s)
	println(&s[0])
	Assign1(s)
	fmt.Println(s) // (1)
	array := [5]int{1, 2, 3, 4, 5}
	Reverse0(array)
	fmt.Println(array) // (2)
	s = []int{1, 2, 3}
	Reverse2(s)
	fmt.Println(s) // (3)
	var a []int
	for i := 1; i <= 3; i++ {
		a = append(a, i)
	}

	fmt.Println(a, len(a), cap(a))
	Reverse2(a)
	fmt.Println(a) // (4)

	var b []int
	for i := 1; i <= 3; i++ {
		b = append(b, i)
	}
	Reverse3(b)
	fmt.Println(b) // (5)
	c := [3]int{1, 2, 3}
	d := c
	c[0] = 999
	fmt.Println(d) // (6) 值拷贝

	println("=====slice的底层是数组指针：https://juejin.cn/post/6844903859257606158====")
	s1 := []int{1, 2, 3} // len=3, cap=3
	a1 := s1
	s1[0] = 888
	s1 = append(s1, 4)

	fmt.Println(a1, len(a1), cap(a1)) // 输出：[888 2 3] 3 3
	fmt.Println(s1, len(s1), cap(s1)) // 输出：[888 2 3 4] 4 6
	println("===== 当append后，slice长度不超过容量cap，新增的元素将直接加在数组中。：https://juejin.cn/post/6844903859257606158====")
	s2 := make([]int, 0, 4)
	s2 = append(s2, 1, 2, 3)
	fmt.Println(s2, len(s2), cap(s2)) // 输出：[1, 2, 3] 3 4
	s2 = append(s2, 4)
	fmt.Println(s2, len(s2), cap(s2)) // 输出：[1, 2, 3] 4 4

	println("=====面试题2：https://segmentfault.com/a/1190000041404278====")
	println("=====1.数组和切片有什么区别？====")
	println("=====2.拷贝大切片一定比拷贝小切片代价大吗？====")

	println("====面试题3：https://juejin.cn/post/6844904177022271501====")
	println("Go Slice探秘——slice作为函数参数传递时，若修改函数中的slice，到底会不会改变原slice的值？")

	println("======场景1===========")
	// myWeight是我7天的体重值序列
	myWeight := []int{11, 12, 13, 14, 15, 16, 17}
	fmt.Printf("myWeight: %v\n", myWeight)
	fmt.Printf("address of myWeight: %p    %p\n", myWeight, &myWeight)
	fmt.Printf("myWeight len: %v, cap: %v\n\n", len(myWeight), cap(myWeight))

	// 重置数据
	resetWeight(myWeight)
	fmt.Printf("myWeight after reset: %v\n", myWeight)
	fmt.Printf("address of myWeight after reset: %p    %p\n", myWeight, &myWeight)
	fmt.Printf("myWeight len: %v, cap: %v\n", len(myWeight), cap(myWeight))

	/* 作者：flappybird
	   链接：https://juejin.cn/post/6844904177022271501
	   来源：稀土掘金
	   著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处 */
	println("======场景2：在函数中向slice添加成员===========")
	myWeight = make([]int, 1, 3)
	fmt.Printf("myWeight: %v\n", myWeight)
	fmt.Printf("address of myWeight: %p       %p\n", myWeight, &myWeight)
	fmt.Printf("myWeight len: %v, cap: %v\n\n", len(myWeight), cap(myWeight))
	addWeightRecord(myWeight)
	fmt.Printf("myWeight after add: %v\n", myWeight)
	fmt.Printf("address of myWeight after add: %p     %p\n", myWeight, &myWeight)
	fmt.Printf("myWeight len: %v, cap: %v\n", len(myWeight), cap(myWeight))

}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func Assign1(s []int) {
	println(&s[0])
	s = []int{6, 6, 6}
	println(&s[0])
}

func Reverse0(s [5]int) {
	for i, j := 0, len(s)-1; i < j; i++ {
		j = len(s) - (i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

func Reverse1(s []int) {
	for i, j := 0, len(s)-1; i < j; i++ {
		j = len(s) - (i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

func Reverse2(s []int) {
	s = append(s, 999)
	for i, j := 0, len(s)-1; i < j; i++ {
		j = len(s) - (i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

func Reverse3(s []int) {
	s = append(s, 999, 1000, 1001)
	fmt.Printf("reverse3: len(s):%d\n", len(s))
	for i, j := 0, len(s)-1; i < j; i++ {
		j = len(s) - (i + 1)
		s[i], s[j] = s[j], s[i]
	}
}
func resetWeight(weight []int) {
	for i := 0; i < len(weight); i++ {
		weight[i] = weight[i] + i*10
	}
	fmt.Printf("address of weight: %p      %p\n\n", weight, &weight)
}

func addWeightRecord(weight []int) {
	weightCap := cap(weight)
	weight[0] = 10
	fmt.Printf("cap of weight: %v\n\n", weightCap)
	for i := 0; i < weightCap-1; i++ {
		weight = append(weight, i)
		fmt.Printf("weight: %v\n", weight)
		fmt.Printf("address of weight: %p     %p\n", weight, &weight)
		fmt.Printf("weight len: %v, cap: %v\n", len(weight), cap(weight))
	}

	fmt.Printf("weight: %v\n", weight)
	fmt.Printf("address of weight: %p     %p\n", weight, &weight)
	fmt.Printf("weight len: %v, cap: %v\n", len(weight), cap(weight))
}
