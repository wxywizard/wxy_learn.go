package exercise

import (
	"fmt"
	"testing"
)

func TestExercise(t *testing.T) {
	fmt.Println("hello world")
	var touHigh int = 6
	var jianHigh int = 25
	var kuang int = 50
	for i := 1; i <= touHigh+jianHigh; i++ {
		for j := 1; j <= kuang; j++ {
			if i <= touHigh {
				if j >= (kuang/2+1)+1-i && j <= (kuang/2+1)-1+i {
					fmt.Print("*")
				} else {
					fmt.Print("-")
				}
			}

			if i > touHigh && i < jianHigh {
				if j >= (kuang/2+1)+2-i && j <= kuang-3*(i-touHigh) {
					fmt.Print("*")
				} else if j < (kuang/2+1)-3+i && j >= 3*(i-touHigh) {
					fmt.Print("*")
				} else {
					fmt.Print("-")
				}
			}
		}
		fmt.Println("")
	}
}
