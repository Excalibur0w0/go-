package main

import (
	"fmt";
	"reflect";
)

func pTypeof(_var interface{}) {
	fmt.Printf("%T", _var)
}

func typeof(_var interface{}) string {
	return reflect.TypeOf(_var).String()
}

func n() {
	fmt.Println()
}

func main() {
	var slice1 []int = make([]int, 10)
	//equal
	slice2 := make([]int, 10)
	fmt.Println("切片数组1:", slice1)
	fmt.Println("切片数组2:", slice2)
	fmt.Println("slice1 的类型为:", typeof(slice1))
	n()

	// 切片拷贝
	var arr []int = []int {1, 3, 5, 3}
	arrSlice1 := arr[:]
	fmt.Println("原始数组", arr)
	fmt.Println("新的切片数组", arrSlice1)
	fmt.Println("从数组三位到末尾开始切片", arr[3:])
	fmt.Println("从数组开始到倒数第三个开始切片", arr[:len(arr) - 3])
	n()

	var arr2 []int = []int {2, 4, 5, 1}
	arr3 := append(arr2, 3)
	fmt.Println("验证切片和原数组的相等性", &arr3 == &arr2)
	n()

	// cap函数用于切片扩容 type, length, capacity
	arrK1 := make([]int, 1, 1)
	arrK1[0] = 1
	arrK2 := make([]int, len(arrK1) * 2, cap(arrK1) * 2)
	// new - src
	copy(arrK2, arrK1)
	fmt.Println("arrK1: ", arrK1, "arrK2", arrK2)
	fmt.Println("arrK1 len: ", len(arrK1), "arrK2 len: ", len(arrK2))
	fmt.Println("arrK1 cap: ", cap(arrK1), "arrK2 cap: ", cap(arrK2))
	n()

}