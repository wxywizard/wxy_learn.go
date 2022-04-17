package main

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"io"
	"learn.go/pkg/apis"
	"log"
	"sync"
)

var _ apis.ChatServiceServer = &chatServer{}

type chatServer struct {
	apis.UnimplementedChatServiceServer
	sync.Mutex
	persons  map[int64]*apis.PersonalInformation
	personCh chan *apis.PersonalInformation
	conn     *gorm.DB
}

func (c *chatServer) RegisterPersons(server apis.ChatService_RegisterPersonsServer) error {
	pis := &apis.PersonalInformationList{}
	for {
		pi, err := server.Recv()
		if err == io.EOF {
			break
		}
		if err != nil && err != io.EOF {
			log.Printf("WARNING: 获取人员注册时失败：%v\n", err)
			return err
		}
		pis.Items = append(pis.Items, pi)
		c.regPerson(pi)
	}
	log.Println("连续得到注册清单：", pis.String())
	return server.SendAndClose(pis)
}

func (c *chatServer) regPerson(pi *apis.PersonalInformation) {
	c.personCh <- pi
}

func (c *chatServer) Register(ctx context.Context, information *apis.PersonalInformation) (*apis.PersonalInformation, error) {
	c.Lock()
	defer c.Unlock()
	c.persons[information.Account] = information
	return information, nil
}

func (c *chatServer) Publish(p *apis.PersonalInformation) error {
	resp := c.conn.Create(p)
	if err := resp.Error; err != nil {
		fmt.Println("入库时失败：", err)
		return err
	}
	fmt.Println("入库成功")
	return nil
}
