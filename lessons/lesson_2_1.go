package lesson

import (
	"fmt"
)

/*
数组
go语言中数组是值类型，数组是一个固定大小的数据集合，数组中的元素类型必须相同。
var arr [10]int
[10]int和[20]int是不同的类型
var arr2 [2]string = ["a", "b"]
go语言一般不使用数组，而使用切片
*/

func Lesson_2_1() {
	fmt.Println("Hello, lesson 2-1!")
	var arr1 [10]int
	fmt.Println("arr1 ---", arr1)
	arr2 := [2]string{"a", "b"}
	fmt.Println("arr2 ---", arr2)
	arr3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println("arr3 ---", arr3)
	var arr4 [2][3]bool
	fmt.Println("arr4 ---", arr4)

	// 数组的遍历
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	} // 1 2 3 4 5

	for i, v := range arr3 {
		fmt.Println(i, v)
	} // 0 1 2 3 4
}
