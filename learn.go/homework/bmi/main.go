package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	for {
		s := process()
		fmt.Println("the physique is " + s)
		fmt.Println("whether to continue？(Y/N)")
		isContinue := read("isContinue")
		upper := strings.ToUpper(strings.TrimSpace(isContinue))
		if "N" == upper {
			os.Exit(-1)
		}
	}
}

func process() string {
	var sex int
	fmt.Println("Please enter sex: ")
	s := read("sex")
	switch strings.TrimSpace(s) {
	case "男":
		sex = man
	case "女":
		sex = woman
	default:
		fmt.Println("the input is invalid")
	}
	fmt.Println("Please enter weight(kg): ")
	w := read("weight")
	wFloat, wErr := strconv.ParseFloat(strings.TrimSpace(w), 64)
	if wErr == nil {
		weight = wFloat
	} else {
		fmt.Println("the input is invalid")
	}
	fmt.Println("Please enter height(m): ")
	h := read("height")
	hFloat, hErr := strconv.ParseFloat(strings.TrimSpace(h), 64)
	if hErr == nil {
		height = hFloat
	} else {
		fmt.Println("the input is invalid")
	}
	fmt.Println("Please enter age: ")
	a := read("age")
	aInt, aErr := strconv.ParseInt(strings.TrimSpace(a), 10, 64)
	if aErr == nil {
		age = aInt
	} else {
		fmt.Println("the input is invalid")
	}
	//计算体脂
	var physiqueResult string
	physique := calculationPhysique(weight, height, age, sex)
	if sex == man {
		physiqueResult = switchMan(physique, age)
	}
	if sex == woman {
		physiqueResult = switchWoman(physique, age)
	}
	return physiqueResult
}

func read(msg string) string {
	reader := bufio.NewReader(os.Stdin)
	readString, err := reader.ReadString('\n')
	if err == nil {
		fmt.Printf("The input "+msg+" was: %s\n", readString)
	}
	return readString
}

/**
BMI=体重（公斤）÷（身高×身高）（米）
体脂率：1.2×BMI+0.23×年龄-5.4-10.8×性别（男为1，女为0）
*/
func calculationPhysique(w float64, h float64, a int64, s int) float64 {
	bmi := w / (h * h)
	result := (1.2*bmi + 0.23*float64(a) - 5.4 - 10.8*float64(s)) / 100
	return result
}

func switchMan(physique float64, age int64) string {
	var result string
	if age >= 18 && age <= 39 {
		if physique >= 0.00 && physique <= 0.10 {
			result = "偏瘦"
		}
		if physique > 0.10 && physique <= 0.16 {
			result = "标注"
		}
		if physique > 0.16 && physique <= 0.21 {
			result = "偏重"
		}
		if physique > 0.21 && physique <= 0.26 {
			result = "肥胖"
		}
		if physique > 0.26 {
			result = "严重肥胖"
		}
	}

	if age >= 40 && age <= 59 {
		if physique >= 0.00 && physique <= 0.11 {
			result = "偏瘦"
		}
		if physique > 0.11 && physique <= 0.17 {
			result = "标注"
		}
		if physique > 0.17 && physique <= 0.22 {
			result = "偏重"
		}
		if physique > 0.22 && physique <= 0.27 {
			result = "肥胖"
		}
		if physique > 0.27 {
			result = "严重肥胖"
		}
	}

	if age >= 60 {
		if physique >= 0.00 && physique <= 0.13 {
			result = "偏瘦"
		}
		if physique > 0.13 && physique <= 0.19 {
			result = "标注"
		}
		if physique > 0.19 && physique <= 0.24 {
			result = "偏重"
		}
		if physique > 0.24 && physique <= 0.29 {
			result = "肥胖"
		}
		if physique > 0.29 {
			result = "严重肥胖"
		}
	}

	return result
}

func switchWoman(physique float64, age int64) string {
	var result string
	if age >= 18 && age <= 39 {
		if physique >= 0.00 && physique <= 0.19 {
			result = "偏瘦"
		}
		if physique > 0.19 && physique <= 0.27 {
			result = "标注"
		}
		if physique > 0.27 && physique <= 0.34 {
			result = "偏重"
		}
		if physique > 0.34 && physique <= 0.39 {
			result = "肥胖"
		}
		if physique > 0.39 {
			result = "严重肥胖"
		}
	}

	if age >= 40 && age <= 59 {
		if physique >= 0.00 && physique <= 0.21 {
			result = "偏瘦"
		}
		if physique > 0.21 && physique <= 0.28 {
			result = "标注"
		}
		if physique > 0.28 && physique <= 0.35 {
			result = "偏重"
		}
		if physique > 0.35 && physique <= 0.40 {
			result = "肥胖"
		}
		if physique > 0.40 {
			result = "严重肥胖"
		}
	}

	if age >= 60 {
		if physique >= 0.00 && physique <= 0.22 {
			result = "偏瘦"
		}
		if physique > 0.22 && physique <= 0.29 {
			result = "标注"
		}
		if physique > 0.29 && physique <= 0.36 {
			result = "偏重"
		}
		if physique > 0.36 && physique <= 0.41 {
			result = "肥胖"
		}
		if physique > 0.41 {
			result = "严重肥胖"
		}
	}

	return result
}
