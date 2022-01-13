package main

import "fmt"

func main() {
	el := Elevator{
		currentFloor: 3,
		totalFloor:   8,
		pressFloors:  []int{4, 5, 2},
		direction:    1,
	}
	err := el.ElevatorRun()
	if err != nil {
		fmt.Println("电梯运行故障")
	}
}
