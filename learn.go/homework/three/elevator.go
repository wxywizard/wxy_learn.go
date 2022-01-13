package main

import (
	"fmt"
)

type Elevator struct {
	//当前楼层
	currentFloor int

	//总楼层
	totalFloor int

	//按下的楼层
	pressFloors []int

	//行进方向0停止，1向上，2向下
	direction int
}

func (el *Elevator) ElevatorRun() {

}

func (el *Elevator) Press(totalFloor int, pressFloor int) (err error) {
	if totalFloor <= 0 || pressFloor <= 0 {
		err = fmt.Errorf("总楼层或者所安楼层不能小于等于0")
		return
	}
	if pressFloor > totalFloor {
		err = fmt.Errorf("按下的楼层不应大于总楼层！")
		return
	}
	el.pressFloors = make([]int, 0)
	el.pressFloors = append(el.pressFloors, pressFloor)
	return nil
}
