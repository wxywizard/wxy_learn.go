package goexample

import "fmt"

func CalcBMI(tall float64, weight float64) (bmi float64, err error) {
	if tall <= 0 {
		err = fmt.Errorf("身高不能是0或者负数")
	}
	if weight <= 0 {
		err = fmt.Errorf("体重不能是0或者负数")
	}
	return weight / (tall * tall), nil
}
