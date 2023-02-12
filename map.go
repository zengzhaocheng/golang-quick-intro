package main

import (
	"fmt"
)

func main() {
	var countryCapitalMap map[string]string /*创建集合 */
	countryCapitalMap = make(map[string]string)
	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India"] = "新德里"
	/*使用键输出地图值 */
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}
	capital, ok := countryCapitalMap["American"] /*如果确定是真实的,则存在,否则不存在 */
	/*fmt.Println(capital) */
	/*fmt.Println(ok) */
	if ok {
		fmt.Println("American 的首都是", capital)
	} else {
		fmt.Println("American 的首都不存在")
	}

	println("======delete======")
	countryCapitalMap = map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}
	fmt.Println("原始地图")
	/* 打印地图 */
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	/*删除元素*/
	delete(countryCapitalMap, "France")
	fmt.Println("法国条目被删除")

	fmt.Println("删除元素后地图")

	/*打印地图*/
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	println("=======面试题1：https://www.jishuchi.com/read/go-interview/5082=====")
	println("1.创建，增删改查：")
	m := map[string]int{}
	// 增加一个key/value对
	m["Tony"] = 10
	// 删除Key Tony
	delete(m, "Tony")
	// 修改Key Tony的值
	m["Tony"] = 20
	// 判断某个Key是否存在
	if _, ok := m["Tony"]; ok {
		fmt.Println("Tony is exists")
	}
	// 遍历map
	for key, value := range m {
		fmt.Printf("Key = %s, Value = %d", key, value)
	}
	// 使用多个值对map进行初始化
	m = map[string]int{
		"Tina":  10,
		"Divad": 20,
		"Tom":   5,
	}

}
