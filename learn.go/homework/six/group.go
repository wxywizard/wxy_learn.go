package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
)

type Group struct {
	conn *gorm.DB
}

func (g *Group) Publish(p *PersonalInformation) error {
	resp := g.conn.Create(p)
	if err := resp.Error; err != nil {
		fmt.Println("创建人***时失败：", err)
		return err
	}
	fmt.Println("创建人***成功")
	return nil
}

func (g *Group) Visit() (data []byte) {
	output := make([]*PersonalInformation, 0)
	resp := g.conn.Where("is_valid = true").Find(&output)
	if resp.Error != nil {
		fmt.Println("查找失败：", resp.Error)
		return
	}

	data, _ = json.Marshal(output)
	fmt.Printf("查询结果：%+v\n", string(data))
	return
}

func (g *Group) delete(p *PersonalInformation) error {
	resp := g.conn.Model(p).Select("id", "is_valid").Updates(p)
	if err := resp.Error; err != nil {
		fmt.Println("更新人状态***时失败：", err)
		return err
	}
	fmt.Println("更新人状态***成功")
	return nil
}
