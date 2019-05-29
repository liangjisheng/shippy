package main

import (
	"fmt"
	"log"
	pb "shippy/user-service/proto/user"

	"github.com/micro/go-micro"
)

func main() {
	db, err := CreateConnection()
	fmt.Printf("%+v\n", db)
	fmt.Printf("err: %v\n", err)
	if err != nil {
		log.Fatalf("connect error: %v\n", err)
	}
	defer db.Close()

	repo := &UserRepository{db}
	// 自动检查 User 结构是否变化
	db.AutoMigrate(&pb.User{})

	s := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)
	// 初始化命令行环境
	// srv.Init() 会加载该微服务的所有配置，比如使用到的插件、设置的环境变量、
	// 命令行参数等，这些配置项会作为微服务的一部分来运行。可使用
	// s.Server().Options() 来获取这些配置
	s.Init()

	// 获取 broker 实例
	// pubSub := s.Server().Options().Broker
	publisher := micro.NewPublisher(topic, s.Client())

	token := &TokenService{repo}
	pb.RegisterUserServiceHandler(s.Server(), &handler{repo, token, publisher})

	if err := s.Run(); err != nil {
		log.Fatalf("user service error: %v\n", err)
	}
}
