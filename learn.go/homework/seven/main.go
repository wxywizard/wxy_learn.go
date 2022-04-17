package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {

}

func connectDb() *gorm.DB {
	conn, err := gorm.Open(mysql.Open("root:root123@tcp(192.168.56.200:3306)/learn_go?parseTime=true"))
	if err != nil {
		log.Fatal("数据库连接失败：", err)
	}
	fmt.Println("连接数据库成功")
	return conn
}
