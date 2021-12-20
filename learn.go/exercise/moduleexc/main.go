package moduleexc

import (
	"fmt"
	gobmi "github.com/wxywizard/go_bmi"
	_ "github.com/wxywizard/go_example/go_bmi"
)

func main() {
	if res, err := gobmi.CalcBMI(1.2, 3.4); err != nil {
		fmt.Println(res)
	}
}
