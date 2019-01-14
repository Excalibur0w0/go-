package main

import "fmt"

type Phone interface {
	call() string;
}

type NokiaPhone struct {
	number int64
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println(nokiaPhone.number)
	fmt.Println("I am Nokia, i can tell you")
}

func main() {
	var phone = new(NokiaPhone)
	
	phone.call()
}