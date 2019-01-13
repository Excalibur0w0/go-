package main

import "fmt"

func showAddress(_var interface{}) {
	fmt.Printf("变量的地址: %x\n", &_var)
	fmt.Println("变量为：", _var)
}

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

func main() {
	var a int = 10
	fmt.Printf("变量的地址: %x\n", &a)
	showAddress(a)
	// 证明go的int类型在go中是值传递
	fmt.Println()

	// 定义指针
	var ip_a *int
	ip_a = &a
	fmt.Println("定义并输出指针存储的变量值", *ip_a)
	fmt.Println()

	// 空指针
	var ptr *int
	fmt.Printf("ptr 的值为: %x\n", ptr)

	// 空指针判断
	if (ptr == nil) {
		fmt.Println("ptr 是空指针")
	} else {
		fmt.Println("ptr 不是空指针")
	}
	fmt.Println()

	// 指针数组
	arr := [...]int{1, 3, 4}
	ptr1Arr := &arr
	var ptr2Arr *[3]int = &arr
	fmt.Println(ptr1Arr)
	fmt.Println(ptr2Arr)
	fmt.Println("ptr1和ptr2比较运算：", ptr1Arr == ptr2Arr)
	fmt.Println("ptrArr的类型", typeof(ptr1Arr))
}