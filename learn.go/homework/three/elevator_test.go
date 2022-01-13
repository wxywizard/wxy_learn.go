package main

import (
	"testing"
)

func TestElevatorRun(t *testing.T) {
	el := Elevator{
		currentFloor: 3,
		totalFloor:   8,
		pressFloors:  []int{4, 5, 2},
		direction:    1,
	}
	err := el.ElevatorRun()
	if err != nil {
		t.Fatal("出现错误，与预期不符")
	}
}

func TestPress(t *testing.T) {
	el := Elevator{}
	totalFloor := 5
	err := el.Press(5, []int{3, 4, 5}, 1)
	t.Logf("当前总楼层是%d", el.totalFloor)
	if err != nil {
		t.Fatal("出现错误，与预期不符")
	}
	if totalFloor != el.totalFloor {
		t.Fatal("出现错误，与预期不符")
	}
}
