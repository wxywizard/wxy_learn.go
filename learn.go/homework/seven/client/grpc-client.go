package main

import (
	"context"
	_ "encoding/json"
	"google.golang.org/grpc/credentials/insecure"
	"learn.go/pkg/apis"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:9010", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c := apis.NewChatServiceClient(conn)
	ret, err := c.Register(context.TODO(), &apis.PersonalInformation{
		Id:       1,
		Nickname: "h1",
		Password: "123",
		Account:  123456,
	})
	if err != nil {
		log.Fatal("注册失败：", err)
	}
	log.Println("注册成功", ret)

	//log.Println("开始批量注册")
	//regCli, err := c.RegisterPersons(context.TODO())
	//if err := regCli.Send(&apis.PersonalInformation{
	//	Name:   "tom-3",
	//	Sex:    "男",
	//	Tall:   1.7,
	//	Weight: 70,
	//	Age:    30,
	//}); err != nil {
	//	log.Fatal("注册时失败：", err)
	//}
	//if err := regCli.Send(&apis.PersonalInformation{
	//	Name:   "tom-1",
	//	Sex:    "男",
	//	Tall:   1.7,
	//	Weight: 70,
	//	Age:    30,
	//}); err != nil {
	//	log.Fatal("注册时失败：", err)
	//}
	//if err := regCli.Send(&apis.PersonalInformation{
	//	Name:   "tom-2",
	//	Sex:    "男",
	//	Tall:   1.7,
	//	Weight: 70,
	//	Age:    30,
	//}); err != nil {
	//	log.Fatal("注册时失败：", err)
	//}
	//regCli.CloseAndRecv()
	//

}
