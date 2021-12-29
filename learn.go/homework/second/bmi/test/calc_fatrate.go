package test

import (
	"fmt"
	"strings"
)

func CalcFatRate(bmi float64, age int, sex string) (fatRate float64, err error) {
	if bmi < 0 {
		err = fmt.Errorf("bmi cannot be negative")
		return
	}
	if age < 0 || age > 150 {
		err = fmt.Errorf("age cannot be negative")
		return
	}

	if strings.Compare("男", sex) != 0 && strings.Compare("女", sex) != 0 {
		err = fmt.Errorf("sex is Illegal")
		return
	}
	sexWeight := 0
	if sex == "男" {
		sexWeight = 1
	} else {
		sexWeight = 0
	}
	fatRate = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexWeight)) / 100
	return
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
