package main

import (
	"fmt"
	bmi "github.com/armstrongli/go-bmi"
)

func main() {
	var resultBMI float64
	if result, err := bmi.BMI(60, 1.83); err == nil {
		fmt.Println(result)
		resultBMI = result
	}
	// 通过本地自定义扩展fatRate
	if fatRate, err := bmi.CalcFatRate(resultBMI, 28, "男"); err == nil {
		fmt.Println(fatRate)
	}
}
