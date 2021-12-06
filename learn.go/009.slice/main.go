package main

import "fmt"

func main() {
	//var a string = "hello"
	//fmt.Println(a)
	//aBytes := []byte(a) //[]int 不能转，[]byte 可以和string互转
	//fmt.Println(aBytes)
	//fmt.Println("修改切片的内容")
	//aBytes[0] = 'H'
	//a = string(aBytes)
	//fmt.Println(a)

	var a string = "你好"
	fmt.Println(a)
	aBytes := []rune(a) //[]int 不能转，[]byte 可以和string互转
	fmt.Println(aBytes)
	fmt.Println("修改切片的内容")
	aBytes[0] = 'H'
	a = string(aBytes)
	fmt.Println(a)

	var s int = 1
	var p = &s
	fmt.Println(&s)
	fmt.Println(*p)
	fmt.Printf("%T", p)

	s2 := []int{1, 2, 3}
	fmt.Println("原始s2切片的数据:", s2)
	fun1(s2)
	fmt.Println("原始s2切片的数据:", s2)

	arr := [3]int{1, 2, 3}
	fmt.Println("原始s2数组的数据:", arr)
	fun2(arr)
	fmt.Println("原始s2数组的数据:", arr)
}

func fun1(s2 []int) {
	fmt.Println("函数中,切片的数据:", s2)
	s2[0] = 100
	fmt.Println("函数中,切片的数据:", s2)
}

func fun2(arr [3]int) {
	fmt.Println("函数中,数组的数据:", arr)
	arr[0] = 100
	fmt.Println("函数中,数组的数据:", arr)
}
