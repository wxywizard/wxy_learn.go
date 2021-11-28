package main

import (
	"fmt"
	"strconv"
)

func main() {
	name := "123"
	atoi, err := strconv.Atoi(name)
	if err == nil {
		fmt.Println(atoi)
	}

	age := 30
	fmt.Printf("%p\n", &age)
	age, time := 32, "时间"
	fmt.Printf("%p\n", &age)
	fmt.Println(age, time)
	{
		age := 3
		fmt.Printf("%p\n", &age)
	}
}
