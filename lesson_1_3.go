package main

import (
	"fmt"
	"math"
)

/*
const 定义 常量
普通枚举
自增枚举: iota

*/

func constants() {
	const (
		a = 1
		c = 3
		d = 4
	)

	const (
		a1 = iota
		b1
		c1
	)

	// iota表示自增值
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	var triangle = int(math.Sqrt(c*c + d*d))

	fmt.Println(a, c, d)
	fmt.Println(a1, b1, c1)
	fmt.Println(b, kb, mb, gb, tb, pb)
	fmt.Println(triangle)
}

func main() {
	constants()
	fmt.Println(math.Pi)
}
