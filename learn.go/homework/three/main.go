package main

func main() {
	el := Elevator{
		currentFloor: 3,
		totalFloor:   8,
		pressFloors:  []int{4, 5, 2},
		direction:    1,
	}
	el.ElevatorRun()
}
