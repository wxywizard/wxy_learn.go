package main

func bubblePerson(arr []*Person) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j].fatRate > arr[j+1].fatRate {
				arr[j].fatRate, arr[j+1].fatRate = arr[j+1].fatRate, arr[j].fatRate
			}
		}
	}
}
