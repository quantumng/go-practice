package main

import "fmt"

/*
变量定义
1.使用var关键字声明变量
2.变量名在前，类型在后
3.可使用var(变量名) = 值 的方式初始化一堆变量
4.可放在函数体内，也可放在函数外（包内）
5.简写方式：变量 := 值
6.go语言具备变量的重复声明，但是不允许变量的重复定义
7.go语言具备类型推导，变量的类型可以不声明
*/
var a int
var b string
var c bool
var d, e, f int
var g, h, i = 1, 2, 3
var (
	a1 int    = 1
	b1 string = "2"
	c1 bool   = true
)

func main() {
	fmt.Println("Hello, world!")
	fmt.Printf("%d %q\n", a, b)
}
