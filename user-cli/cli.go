package main

import (
	"context"
	"log"
	"os"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"

	pb "shippy/user-service/proto/user"
)

func main() {
	cmd.Init()

	service := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
		micro.Flags(
			cli.StringFlag{
				Name:  "name",
				Usage: "You full name",
			},
			cli.StringFlag{
				Name:  "email",
				Usage: "Your email",
			},
			cli.StringFlag{
				Name:  "password",
				Usage: "Your password",
			},
			cli.StringFlag{
				Name:  "company",
				Usage: "Your company",
			},
		),
	)

	client := pb.NewUserService("go.micro.srv.user", service.Client())

	service.Init(
		micro.Action(func(c *cli.Context) {
			name := c.String("name")
			email := c.String("email")
			password := c.String("password")
			company := c.String("company")

			name = "liangjisheng"
			email = "1294851990@qq.com"
			password = "liangjisheng"
			company = "liangjisheng"

			r, err := client.Create(context.TODO(), &pb.User{
				Name:     name,
				Email:    email,
				Password: password,
				Company:  company,
			})
			if err != nil {
				log.Fatalf("Could not create: %v", err)
			}
			log.Printf("Created: %v", r.User.Id)

			getAll, err := client.GetAll(context.Background(), &pb.Request{})
			if err != nil {
				log.Fatalf("Could not list users: %v", err)
			}
			for _, v := range getAll.Users {
				log.Println(v)
			}

			authResp, err := client.Auth(context.Background(), &pb.User{
				Email:    email,
				Password: password,
			})

			if err != nil {
				log.Printf("auth failed: %v", err)
			}
			log.Println("token:", authResp.Token)

			os.Exit(0)
		}),
	)

	// 启动客户端
	if err := service.Run(); err != nil {
		log.Println(err)
	}
}

func client1() {
	cmd.Init()
	// 创建 user-service 微服务的客户端
	client := pb.NewUserService("go.micro.srv.user", microclient.DefaultClient)

	// 暂时将用户信息写死在代码中
	name := "Ewan Valentine"
	email := "ewan.valentine89@gmail.com"
	password := "test123"
	company := "BBC"

	resp, err := client.Create(context.TODO(), &pb.User{
		Name:     name,
		Email:    email,
		Password: password,
		Company:  company,
	})
	if err != nil {
		log.Fatalf("call Create error: %v", err)
	}
	log.Println("created: ", resp.User.Id)

	// allResp, err := client.GetAll(context.Background(), &pb.Request{})
	// if err != nil {
	// 	log.Fatalf("call GetAll error: %v", err)
	// }
	// for i, u := range allResp.Users {
	// 	log.Printf("user_%d: %v\n", i, u)
	// }

	authResp, err := client.Auth(context.TODO(), &pb.User{
		Email:    email,
		Password: password,
	})
	if err != nil {
		log.Fatalf("auth failed: %v", err)
	}
	log.Println("token: ", authResp.Token)

	// 直接退出即可
	os.Exit(0)
}
