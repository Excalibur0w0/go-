package main

import "fmt"
// import "time"

// golang不支持默认值
type Book struct {
	name string
	value float32
}

func main() {
	var book Book

	fmt.Println(book)
	fmt.Printf("%T", book)
}