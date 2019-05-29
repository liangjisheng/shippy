package main

import (
	"context"
	"log"

	userPb "shippy/user-service/proto/user"

	"github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/broker/nats"
)

const topic = "user.created"

// Subscriber ...
type Subscriber struct{}

func main() {
	srv := micro.NewService(
		micro.Name("go.micro.srv.email"),
		micro.Version("latest"),
	)
	// 初始化命令行环境
	srv.Init()

	// pubSub := srv.Server().Options().Broker
	// if err := pubSub.Connect(); err != nil {
	// 	log.Fatalf("broker connect error: %v\n", err)
	// }

	// // 订阅消息
	// _, err := pubSub.Subscribe(topic, func(pub broker.Publication) error {
	// 	var user *userPb.User
	// 	if err := json.Unmarshal(pub.Message().Body, &user); err != nil {
	// 		return err
	// 	}
	// 	log.Printf("[Create User]: %v\n", user)
	// 	go sendEmail(user)
	// 	return nil
	// })

	// if err != nil {
	// 	log.Printf("sub error: %v\n", err)
	// }

	micro.RegisterSubscriber(topic, srv.Server(), new(Subscriber))

	if err := srv.Run(); err != nil {
		log.Fatalf("srv run error: %v\n", err)
	}
}

// Process ...
func (sub *Subscriber) Process(ctx context.Context, user *userPb.User) error {
	log.Println("[Picked up a new message]")
	log.Println("[Sending email to]:", user.Name)
	return nil
}

func sendEmail(user *userPb.User) error {
	log.Printf("[SENDING A EMAIL TO %s...]", user.Name)
	return nil
}
