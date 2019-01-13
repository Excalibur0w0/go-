package main

import "fmt"

func main() {
	// 初始化长度为5的float32数组
	arr1Float32 := [5]float32{.9, .1, .3, 1.32}
	// 不指定开始数组长度，根据内容初始化
	arr2Float32 := [...]float32{1.12, 33, 12.3}
	// 默认初始化
	var arr3Float64 [10] float64
	// or
	arr4Float64 := [10]float64{}

	fmt.Println(arr1Float32)
	fmt.Println(arr2Float32)
	fmt.Println(arr3Float64)
	fmt.Println(arr4Float64)
}