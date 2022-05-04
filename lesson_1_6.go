package main

import (
	"fmt"
)

/*
func eval(a, b *int) {}

eval(&a, &b)

go语言中，函数参数是按值传递的，即传递的是参数的副本，而不是参数本身。
go语言指针是指向一个变量的指针，指针变量的值是指针的地址。
go语言指针是不能运算的，但是可以赋值。

*/

func main() {
	fmt.Println("Hello, World!")
	var a int = 1
	var pa *int = &a
	*pa = 2
	fmt.Println(a)

}
