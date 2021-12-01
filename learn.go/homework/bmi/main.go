package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	man   int = 1
	woman int = 0
)

var (
	weight float64
	height float64
	age    int64
)

func main() {
	var sex int
	fmt.Println("Please enter sex: ")
	s := read("sex")
	switch s {
	case "男":
		sex = 1
	case "女":
		sex = 0
	default:
		fmt.Println("the input is invalid")
	}
	fmt.Println("Please enter weight(kg): ")
	w := read("weight")
	wFloat, wErr := strconv.ParseFloat(w, 64)
	if wErr == nil {
		weight = wFloat
	} else {
		fmt.Println("the input is invalid")
	}
	fmt.Println("Please enter height(m): ")
	h := read("height")
	hFloat, hErr := strconv.ParseFloat(h, 64)
	if hErr == nil {
		height = hFloat
	} else {
		fmt.Println("the input is invalid")
	}
	fmt.Println("Please enter age: ")
	a := read("age")
	aInt, aErr := strconv.ParseInt(a, 10, 64)
	if aErr == nil {
		age = aInt
	} else {
		fmt.Println("the input is invalid")
	}
	//计算体脂
	physique := calculationPhysique(weight, height, age, sex)
	fmt.Sprintf("%.2f", physique)
	//printf, wErr := fmt.Printf("%%", res)
}

func read(msg string) string {
	reader := bufio.NewReader(os.Stdin)
	readString, err := reader.ReadString('\n')
	if err == nil {
		fmt.Printf("The input "+msg+"was: %s\n", readString)
	}
	return readString
}

/**
BMI=体重（公斤）÷（身高×身高）（米）
体脂率：1.2×BMI+0.23×年龄-5.4-10.8×性别（男为1，女为0）
*/
func calculationPhysique(w float64, h float64, a int64, s int) float64 {
	bmi := w / (h * h)
	result := 1.2*bmi + 0.23*(float64)(a) - 5.4 - 10.8*(float64)(s)
	return result
}
