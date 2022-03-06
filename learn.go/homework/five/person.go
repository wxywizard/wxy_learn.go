package main

import (
	"fmt"
	"os"
)

type Persons []*Person

type Person struct {
	name    string
	fatRate float64
}

/**
注册
*/
func (p *Person) register(name string, fatRate float64) {
	p.name = name
	p.fatRate = fatRate
}

func (p *Person) writeFileWithAppendJson(data []byte, filePath string) error {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777) // linux file permission settings
	if err != nil {
		fmt.Println("无法打开文件", filePath, "错误信息是：", err)
		os.Exit(1)
	}
	defer file.Close()

	_, err = file.Write(append(data, '\n'))
	if err != nil {
		return err
	}
	return nil
}
