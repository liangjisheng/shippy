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
	// cmd.Init()

	client := pb.NewShippingService("go.micro.srv.consignment", service.Client())
	// client := pb.NewShippingService("go.micro.srv.consignment", microclient.DefaultClient)

	// 在命令行中指定新的货物信息 json 文件
	infoFile := DEFAULT_INFO_FILE
	if len(os.Args) > 1 {
		infoFile = os.Args[1]
	}

	consignment, err := parseFile(infoFile)
	if err != nil {
		log.Fatalf("parse info file error: %v\n", err)
	}

	resp, err := client.CreateConsignment(context.TODO(), consignment)
	if err != nil {
		log.Fatalf("create consignment error: %v\n", err)
	}
	log.Printf("created: %t\n", resp.Created)

	// resp, err = client.GetConsignments(context.Background(), &pb.GetRequest{})
	// if err != nil {
	// 	log.Fatalf("failed to list consignments: %v\n", err)
	// }

	// for _, c := range resp.Consignments {
	// 	log.Printf("%+v", c)
	// }
}
