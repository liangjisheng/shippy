package main

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"

	pb "shippy/consignment-service/proto/consignment"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"
)

const (
	// ADDRESS           = "localhost:50051"
	DEFAULT_INFO_FILE = "consignment.json"
)

// 读取 consignment.json 中记录的货物信息
func parseFile(fileName string) (*pb.Consignment, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var consignment *pb.Consignment
	err = json.Unmarshal(data, &consignment)
	if err != nil {
		return nil, errors.New("consignment.json file content error")
	}
	return consignment, nil
}

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	// Init will parse the command line flags.
	service.Init()

	client := pb.NewShippingService("go.micro.srv.consignment", service.Client())

	// 在命令行中指定新的货物信息 json 文件
	if len(os.Args) < 3 {
		log.Fatalln("Not enough arguments, expecing file and token.")
	}
	infoFile := os.Args[1]
	token := os.Args[2]

	log.Println("infoFile:", infoFile)
	log.Println("token:", token)

	consignment, err := parseFile(infoFile)
	if err != nil {
		log.Fatalf("parse info file error: %v\n", err)
	}

	// 创建带有用户 token 的 context
	// consignment-service 服务端将从中取出 token，解密取出用户身份
	tokenContext := metadata.NewContext(context.Background(), map[string]string{
		"token": token,
	})

	// 调用 RPC
	// 将货物存储到指定用户的仓库里
	resp, err := client.CreateConsignment(tokenContext, consignment)
	if err != nil {
		log.Fatalf("create consignment error: %v\n", err)
	}
	log.Printf("created: %t\n", resp.Created)

	// 列出目前所有托运的货物
	resp, err = client.GetConsignments(tokenContext, &pb.GetRequest{})
	if err != nil {
		log.Fatalf("failed to list consignments: %v\n", err)
	}

	for i, c := range resp.Consignments {
		log.Printf("consignment_%d: %v\n", i, c)
	}
}
