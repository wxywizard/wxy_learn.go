package main

import (
	"bufio"
	"fmt"
	"github.com/shopspring/decimal"
	"os"
	"strconv"
	"strings"
)

const (
	man   int    = 1
	woman int    = 0
	exit  string = "N"
)

var (
	weight        float64
	height        float64
	age           int64
	count         int64
	physiqueCount float64
)

func main() {
	for {
		result, physique := process()
		fmt.Println(result)
		//输出总人数和平均体脂率
		count++
		physiqueCount += physique
		avg := physiqueCount / float64(count)
		round := decimal.NewFromFloat(avg).Round(2).String()
		formatInt := decimal.NewFromInt(count).String()
		//formatInt := strconv.FormatInt(count, 36)
		//value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", avg), 64)
		//formatAvg:= strconv.FormatFloat(value, 'E', -1, 64)
		fmt.Println("当前输入总人数为：" + formatInt + " 平均体脂率为：" + round)
		//是否退出程序
		fmt.Println("whether to continue？(Y/N)")
		isContinue := read("isContinue")
		upper := strings.ToUpper(strings.TrimSpace(isContinue))
		if exit == upper {
			os.Exit(-1)
		}
	}
}

func process() (string, float64) {
	fmt.Print("Please enter your name: ")
	name := read("height")
	var sex int
	fmt.Print("Please enter sex: ")
	s := read("sex")
	switch s {
	case "男":
		sex = man
	case "女":
		sex = woman
	default:
		fmt.Println("the input is invalid")
	}
	fmt.Print("Please enter weight(kg): ")
	w := read("weight")
	wFloat, wErr := strconv.ParseFloat(w, 64)
	if wErr == nil {
		weight = wFloat
	} else {
		fmt.Println("the input is invalid")
	}
	fmt.Print("Please enter height(m): ")
	h := read("height")
	hFloat, hErr := strconv.ParseFloat(h, 64)
	if hErr == nil {
		height = hFloat
	} else {
		fmt.Println("the input is invalid")
	}
	fmt.Print("Please enter age: ")
	a := read("age")
	aInt, aErr := strconv.ParseInt(a, 10, 64)
	if aErr == nil {
		age = aInt
	} else {
		fmt.Println("the input is invalid")
	}
	//计算体脂
	var physiqueResult string
	bmi, physique := calculationPhysique(weight, height, age, sex)
	if sex == man {
		physiqueResult = switchMan(physique, age)
	}
	if sex == woman {
		physiqueResult = switchWoman(physique, age)
	}
	strBmi := decimal.NewFromFloat(bmi).Round(2).String()
	//strBmi:= strconv.FormatFloat(bmi, 'E', -1, 64)
	result := name + "：BMI为：" + strBmi + " 体脂" + physiqueResult
	return result, physique
}

func read(msg string) string {
	reader := bufio.NewReader(os.Stdin)
	readString, err := reader.ReadString('\n')
	if err == nil {
		fmt.Printf("The input "+msg+" was: %s\n", readString)
	}
	return strings.TrimSpace(readString)
}

/**
BMI=体重（公斤）÷（身高×身高）（米）
体脂率：1.2×BMI+0.23×年龄-5.4-10.8×性别（男为1，女为0）
*/
func calculationPhysique(w float64, h float64, a int64, s int) (float64, float64) {
	bmi := w / (h * h)
	result := (1.2*bmi + 0.23*float64(a) - 5.4 - 10.8*float64(s)) / 100
	return bmi, result
}

func switchMan(physique float64, age int64) string {
	var result string
	if age >= 18 && age <= 39 {
		if physique >= 0.00 && physique <= 0.10 {
			result = "目前是：偏瘦。要多吃多锻炼，增强体质"
		}
		if physique > 0.10 && physique <= 0.16 {
			result = "目前是：标准。太棒了，要保持哦"
		}
		if physique > 0.16 && physique <= 0.21 {
			result = "目前是：偏重。吃完饭多散散步，消化消化"
		}
		if physique > 0.21 && physique <= 0.26 {
			result = "目前是：肥胖。少吃点，多运动运动"
		}
		if physique > 0.26 {
			result = "目前是：严重肥胖。健身游泳跑步，即刻开始"
		}
	}

	if age >= 40 && age <= 59 {
		if physique >= 0.00 && physique <= 0.11 {
			result = "目前是：偏瘦。要多吃多锻炼，增强体质"
		}
		if physique > 0.11 && physique <= 0.17 {
			result = "目前是：标准。太棒了，要保持哦"
		}
		if physique > 0.17 && physique <= 0.22 {
			result = "目前是：偏重。吃完饭多散散步，消化消化"
		}
		if physique > 0.22 && physique <= 0.27 {
			result = "目前是：肥胖。少吃点，多运动运动"
		}
		if physique > 0.27 {
			result = "目前是：严重肥胖。健身游泳跑步，即刻开始"
		}
	}

	if age >= 60 {
		if physique >= 0.00 && physique <= 0.13 {
			result = "目前是：偏瘦。要多吃多锻炼，增强体质"
		}
		if physique > 0.13 && physique <= 0.19 {
			result = "目前是：标准。太棒了，要保持哦"
		}
		if physique > 0.19 && physique <= 0.24 {
			result = "目前是：偏重。吃完饭多散散步，消化消化"
		}
		if physique > 0.24 && physique <= 0.29 {
			result = "目前是：肥胖。少吃点，多运动运动"
		}
		if physique > 0.29 {
			result = "目前是：严重肥胖。健身游泳跑步，即刻开始"
		}
	}

	return result
}

func switchWoman(physique float64, age int64) string {
	var result string
	if age >= 18 && age <= 39 {
		if physique >= 0.00 && physique <= 0.19 {
			result = "目前是：偏瘦。要多吃多锻炼，增强体质"
		}
		if physique > 0.19 && physique <= 0.27 {
			result = "目前是：标准。太棒了，要保持哦"
		}
		if physique > 0.27 && physique <= 0.34 {
			result = "目前是：偏重。吃完饭多散散步，消化消化"
		}
		if physique > 0.34 && physique <= 0.39 {
			result = "目前是：肥胖。少吃点，多运动运动"
		}
		if physique > 0.39 {
			result = "目前是：严重肥胖。健身游泳跑步，即刻开始"
		}
	}

	if age >= 40 && age <= 59 {
		if physique >= 0.00 && physique <= 0.21 {
			result = "目前是：偏瘦。要多吃多锻炼，增强体质"
		}
		if physique > 0.21 && physique <= 0.28 {
			result = "目前是：标准。太棒了，要保持哦"
		}
		if physique > 0.28 && physique <= 0.35 {
			result = "目前是：偏重。吃完饭多散散步，消化消化"
		}
		if physique > 0.35 && physique <= 0.40 {
			result = "目前是：肥胖。少吃点，多运动运动"
		}
		if physique > 0.40 {
			result = "目前是：严重肥胖。健身游泳跑步，即刻开始"
		}
	}

	if age >= 60 {
		if physique >= 0.00 && physique <= 0.22 {
			result = "目前是：偏瘦。要多吃多锻炼，增强体质"
		}
		if physique > 0.22 && physique <= 0.29 {
			result = "目前是：标准。太棒了，要保持哦"
		}
		if physique > 0.29 && physique <= 0.36 {
			result = "目前是：偏重。吃完饭多散散步，消化消化"
		}
		if physique > 0.36 && physique <= 0.41 {
			result = "目前是：肥胖。少吃点，多运动运动"
		}
		if physique > 0.41 {
			result = "目前是：严重肥胖。健身游泳跑步，即刻开始"
		}
	}

	return result
}
