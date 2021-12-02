package main

import "fmt"

func main() {
	//line1 := [4]float64{1.2, 2.4, 6.8, 7.9}
	//line2 := [4]float64{1.2, 2.4, 6.8, 7.6}
	line1 := InputLine("line1")
	line2 := InputLine("line2")
	result := parallel(line1[0], line1[1], line2[0], line2[1], line1[2], line1[3], line2[2], line2[3])
	fmt.Println(result)
}

func parallel(x1 float64, y1 float64, x2 float64, y2 float64,
	x3 float64, y3 float64, x4 float64, y4 float64) string {
	var result string
	if x1 == x2 || x3 == x4 {
		if x1 == x2 && x3 == x4 {
			result = "两直线平行"
		}
	} else {
		p1 := (y2 - y1) / (x2 - x1)
		p2 := (y4 - y3) / (x4 - x3)
		if p1 == p2 {
			result = "两直线平行"
		}
	}
	return result
}

func InputLine(msg string) []float64 {
	length := 0
	fmt.Print("Enter the number of floats for " + msg + ": ")
	fmt.Scanln(&length)
	input := make([]float64, length)
	for i := 0; i < length; i++ {
		fmt.Print("Enter a float: ")
		fmt.Scanf("%f", &input[i])
	}
	return input
}
