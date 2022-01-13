package main

import (
	"fmt"
	"sort"
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

func (el *Elevator) ElevatorRun() (err error) {
	if len(el.pressFloors) == 0 {
		err = fmt.Errorf("当前没有人按电梯！")
		return
	}
	upFloor := make([]int, 0)
	downFloor := make([]int, 0)
	//将电梯按键区分为上行和下行
	for _, item := range el.pressFloors {
		if item > el.currentFloor {
			upFloor = append(upFloor, item)
		} else if item < el.currentFloor {
			downFloor = append(downFloor, item)
		}
	}
	sort.Ints(upFloor)
	sort.Sort(sort.Reverse(sort.IntSlice(downFloor)))
	//判断电梯行进方向
	if el.direction == 0 {
		if el.pressFloors[0] > el.currentFloor {
			el.direction = 1
		} else if el.pressFloors[0] < el.currentFloor {
			el.direction = 2
		}
	}

	if el.direction == 1 {
		upFloor, downFloor = el.upFloorDirection(upFloor, downFloor)
	}

	if el.direction == 2 {
		el.downFloorDirection(downFloor, upFloor)
	}
	return nil
}

func (el *Elevator) downFloorDirection(downFloor []int, upFloor []int) {
	if len(downFloor) <= 0 {
		return
	}
	for _, item := range downFloor {
		el.currentFloor = item
		fmt.Println("当前电梯下行到达", el.currentFloor, "层")
	}
	start := downFloor[len(downFloor)-1]
	end := el.totalFloor

	if len(upFloor) > 0 {
		el.direction = 1
		end = upFloor[len(upFloor)-1]
		var temp []int
		for i := start; i <= end; i++ {
			temp = append(temp, i)
		}
		for _, item := range temp {
			el.currentFloor = item
			fmt.Println("当前电梯上行到达", el.currentFloor, "层")
		}
		upFloor = nil
		temp = nil
	}
	downFloor = nil
}

func (el *Elevator) upFloorDirection(upFloor []int, downFloor []int) ([]int, []int) {
	for _, item := range upFloor {
		el.currentFloor = item
		fmt.Println("当前电梯上行到达", el.currentFloor, "层")
	}
	end := upFloor[len(upFloor)-1]
	start := 1
	if len(downFloor) > 0 {
		start = downFloor[len(downFloor)-1]
		el.direction = 2
		var temp []int
		for i := end; i >= start; i-- {
			temp = append(temp, i)
		}
		for _, item := range temp {
			el.currentFloor = item
			fmt.Println("当前电梯下行到达", el.currentFloor, "层")
		}
		downFloor = nil
		temp = nil
	}
	upFloor = nil
	return upFloor, downFloor
}

func (el *Elevator) Press(totalFloor int, pressFloor []int, direction int) (err error) {
	if totalFloor <= 0 {
		err = fmt.Errorf("总楼层不能小于等于0")
		return
	}
	for _, item := range pressFloor {
		if item > totalFloor {
			err = fmt.Errorf("按下的楼层不应大于总楼层！")
			return
		}
		if item == el.currentFloor {
			err = fmt.Errorf("电梯当前已在%d层", item)
			return
		}
	}
	el.pressFloors = make([]int, 0)
	el.pressFloors = append(el.pressFloors, pressFloor...)
	el.direction = direction
	el.totalFloor = totalFloor
	return nil
}
