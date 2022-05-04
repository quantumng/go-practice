package lesson

import (
	"fmt"
	"io/ioutil"
)

/*
for, if 后面的条件没有括号
if 条件里也可以定义变量(变量作用域为if条件的块)
没有while （使用for替代）
for 可以没有起始赋值语句，没有递增语句，甚至可以没有结束语句（相当于无限循环）
switch 后面的条件没有括号, 不需要break
不需要break的话，可以使用fallthrough, 会继续执行下一个case
可以直接switch 多个case
*/

func returnBool(v int) bool {
	if v > 0 {
		return true
	}
	return false
}

func returnValue(v int) string {
	res := ""
	if result := ""; v > 0 {
		res = "positive"
		result = "positive"
		fmt.Println(result)
	} else if v < 0 {
		res = "negative"
	} else {
		res = "zero"
	}
	return res
}

func readFile() {
	const filename = "test.txt"
	file, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	} else {
		fmt.Printf("%s \n", string(file))
	}
}

func switchCase(grade int) string {
	result := ""
	switch {
	case grade < 60:
		result = "D"
		// fallthrough
	case grade < 70:
		result = "C"
	case grade < 80:
		result = "B"
	case grade <= 100:
		result = "A"
	default:
		result = "error"
	}
	return result
}

func countdown(n int) {
	for i := n; i >= 0; i-- {
		fmt.Println(i)
	}
}

func Lesson_1_4() {
	fmt.Println(returnBool(1))
	fmt.Println(returnValue(5))
	fmt.Println(returnValue(-5))
	fmt.Println(returnValue(0))
	readFile()

	fmt.Println(switchCase(50))
	fmt.Println(switchCase(90))
	countdown(3)
}
