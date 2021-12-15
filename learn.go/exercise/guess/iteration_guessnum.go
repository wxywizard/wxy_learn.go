package guess

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Guess(left, right uint) {
	guessed := (left + right) / 2
	var getFormInput string
	fmt.Println("我猜是：", guessed)
	fmt.Print("如果高了，输入1，如果低了，输入0；对了，输入9：")
	getFormInput = read()
	switch getFormInput {
	case "1":
		if left == right {
			fmt.Println("你是不是改主意了？")
			return
		}
		Guess(left, guessed-1)
	case "0":
		if left == right {
			fmt.Println("你是不是改主意了？")
			return
		}
		Guess(guessed-1, right)
	case "9":
		fmt.Println("你心里的数字是：", guessed)

	}
}

func read() (str string) {
	reader := bufio.NewReader(os.Stdin)
	if readString, err := reader.ReadString('\n'); err == nil {
		str = strings.TrimSpace(readString)
	}
	return str
}
