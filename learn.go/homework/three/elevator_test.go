package main

import (
	"fmt"
	"testing"
)

func TestRunCase1(t *testing.T) {
	el := Elevator{
		currentFloor: 1,
		totalFloor:   5,
		pressFloors:  make([]int, 0),
		direction:    0,
	}
	expectCurrentFloor := 1
	t.Logf("当前总楼层是%d", el.totalFloor)
	if len(el.pressFloors) == 0 {
		if el.currentFloor != expectCurrentFloor {
			t.Fatal("没有人按电梯，应该停止不动")
		}
	}
}

func TestRunCase2(t *testing.T) {
	el := Elevator{
		currentFloor: 1,
		totalFloor:   5,
		pressFloors:  []int{3},
		direction:    0,
	}
	//expectCurrentFloor := 3
	t.Logf("当前总楼层是%d", el.totalFloor)
	for _, item := range el.pressFloors {
		fmt.Println(item)
	}
}
