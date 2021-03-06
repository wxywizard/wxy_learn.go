package main

import "fmt"

type Test struct {
	name string
}

/*
func (t Test) Close() {
	fmt.Println(t.name, "closed")
}
*/

func (t *Test) Close() {
	fmt.Println(t.name, "closed")
}

func main() {
	ts := []Test{{"a"}, {"b"}, {"c"}}
	for _, t := range ts {
		fmt.Println(t)
		//defer func(t Test){
		//	t.Close()
		//}(t)
		defer t.Close()
	}
}
