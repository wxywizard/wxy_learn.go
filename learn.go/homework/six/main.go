package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

func main() {
	db := connectDb()
	g := &Group{
		conn: db,
	}
	pi := &PersonalInformation{
		Id:      3,
		Name:    "小虎",
		Sex:     "男",
		Tall:    1.80,
		Weight:  60,
		Age:     28,
		Time:    time.Now(),
		IsValid: true,
	}
	err := g.Publish(pi)
	if err != nil {
		return
	}
	dpi := &PersonalInformation{
		Id:      2,
		IsValid: false,
	}
	derr := g.delete(dpi)
	if derr != nil {
		return
	}
	g.Visit()

}

func connectDb() *gorm.DB {
	conn, err := gorm.Open(mysql.Open("root:root123@tcp(192.168.56.200:3306)/learn_go?parseTime=true"))
	if err != nil {
		log.Fatal("数据库连接失败：", err)
	}
	fmt.Println("连接数据库成功")
	return conn
}
