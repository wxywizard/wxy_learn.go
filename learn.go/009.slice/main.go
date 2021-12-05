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

}
