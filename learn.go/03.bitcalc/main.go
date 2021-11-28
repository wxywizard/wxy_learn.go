package main

import "fmt"

func main() {
	//	a, b := 100, 31
	//	fmt.Println(a ^ b)
	//	fmt.Println(b ^ a)
	//	arr := []int{4, 3, 4, 5, 6, 7, 3, 5, 6}
	//	result := -1
	//	for _, item := range arr {
	//		if result < 0 {
	//			result = item
	//		} else {
	//			result = result ^ item
	//		}
	//	}
	//	fmt.Println(result)

	var money int = 20
	//var busy bool = true
	switch money {
	case 20:
		fmt.Println("456")
		fmt.Println("123")
		fallthrough
	case 200:
		fmt.Println("12333")
	}

	var num interface{} = 100
	switch newNum := num.(type) {
	case int:
		fmt.Println("num是int", newNum+1.0)
	case int64:
		fmt.Println("num是int64", newNum+3.0)
	case float32:
		fmt.Println("newNum是float32", newNum+1.32)
	case float64:
		fmt.Println("newNum是float64", newNum+3.655)
	default:
		fmt.Println("未知类型")
	}

	switch num.(type) {
	case int, int64, int32, int16:
		fmt.Println("num是整数")
	case float32:
		fmt.Println("newNum是float32")
	case float64:
		fmt.Println("newNum是float64")
	default:
		fmt.Println("未知类型")
	}
}
