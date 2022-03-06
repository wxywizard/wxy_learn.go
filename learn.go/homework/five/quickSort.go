package main

func quickSortPerson(arr []*Person, start, end int) {
	pivotIdx := (start + end) / 2
	pivotV := arr[pivotIdx].fatRate
	l, r := start, end
	for l <= r {
		for arr[l].fatRate < pivotV {
			l++
		}
		for arr[r].fatRate > pivotV {
			r--
		}
		if l >= r {
			break
		}

		arr[l].fatRate, arr[r].fatRate = arr[r].fatRate, arr[l].fatRate
		l++
		r--
	}
	if l == r {
		l++
		r--
	}
	if r > start {
		quickSortPerson(arr, start, r)
	}
	if l < end {
		quickSortPerson(arr, l, end)
	}
}
