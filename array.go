package main

import "fmt"

func ArrayAppend() []int {
	arr := make([]int, 5)
	fmt.Printf("arr.len: %d; arr.cap : %d \n", len(arr), cap(arr))
	arr = append(arr, 10)
	fmt.Printf("arr.len: %d; arr.cap: %d \n", len(arr), cap(arr))
	return arr
}

// 向函数传递数组

func getAverage(arr []float32, size int) float32 {
	var i int
	var avg, sum float32
	for i = 0; i < size; i++ {
		sum += float32(arr[i])
	}
	avg = sum / float32(size)
	return avg
}
func main() {
	var n [10]int
	var i, j int
	for i = 0; i < 10; i++ {
		n[i] = i + 100 /* 设置匀速为 i+ 100*/
	}
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d \n", j, n[j])
	}

	var k int
	balance := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	for i = 0; i < 5; i++ {
		fmt.Printf("balance[%d] = %f \n", i, balance[i])
	}
	balance2 := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	for j = 0; j < 5; j++ {
		fmt.Printf("balance2[%d] = %f\n", j, balance2[j])
	}
	balance3 := [5]float32{1: 2.0, 3: 7.0}
	for k = 0; k < 5; k++ {
		fmt.Printf("balance3[%d] = %f\n", k, balance3[k])
	}

	println("=====数组面试题：https://learnku.com/articles/27630")
	var test []int = ArrayAppend()
	println(test)

	println("====多维数组=====")
	values := [][]int{}
	row1 := []int{1, 2, 3}
	row2 := []int{4, 5, 6}
	row3 := []int{7, 8, 9, 0}
	values = append(values, row1)
	values = append(values, row2)
	values = append(values, row3)

	fmt.Println("row 1")
	fmt.Println(values[0])
	fmt.Println("row 2")
	fmt.Println(values[1])
	fmt.Println("row 3")
	fmt.Println(values[2])

	fmt.Println("第一个元素：")
	fmt.Println(values[0][0])

	println("======向函数传递数组=========")
	balance5 := []float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	avg1 := getAverage(balance5, 5)
	fmt.Printf("平均值：%f\n", avg1)
}
