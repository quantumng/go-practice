package lesson

import (
	"fmt"
)

/*
函数
func eval(a, b int, op string) int
返回值类型写在最后面
可以返回多个值
函数作为参数
没有默认参数，可选参数
可变参数列表
*/

func operation(a, b int, op string) int {
	result := 0
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unsupported operation: " + op)

	}
	return result
}

func makeSplitFunc(a, b int) (r, q int) {
	return a / b, a % b
}

func funcArgs(f func(a, b int) (r, q int), c, d int, op string) int {
	one, two := f(c, d)
	return operation(one, two, op)
}

func Lesson_1_5() {

	fmt.Println(operation(1, 2, "+"))

	fmt.Println(funcArgs(makeSplitFunc, 16, 5, "+"))

}
