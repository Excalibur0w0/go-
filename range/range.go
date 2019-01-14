package main 

import "fmt"

func main() {
	nums := []int{2, 3, 5}
	sum := 0   
	// 解构 
	for _, num := range nums {
			sum += num
	}
	fmt.Println("sum:", sum)
	fmt.Println("nums:", nums)
	fmt.Println()

	
	//range也可以用在map的键值对上。
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
			fmt.Printf("%s -> %s\n", k, v)
	}
	//range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
	for i, c := range "go" {
			fmt.Println(i, c)
	}
}