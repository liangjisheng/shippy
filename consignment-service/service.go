package main

import (
	"context"
	"log"

	pb "shippy/consignment-service/proto/consignment"

	"github.com/micro/go-micro"
)

// IRepository 仓库接口
type IRepository interface {
	Create(consignment *pb.Consignment) (*pb.Consignment, error) // 存放新货物
	GetAll() []*pb.Consignment
}

// Repository 我们存放多批货物的仓库，实现了 IRepository 接口
type Repository struct {
	consignments []*pb.Consignment
}

// Create ...
func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.consignments = append(repo.consignments, consignment)
	return consignment, nil
}

// GetAll ...
func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

// 定义微服务
type service struct {
	repo Repository
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, rsp *pb.Response) error {
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}
	// resp = &pb.Response{Created: true, Consignment: consignment}
	// fmt.Printf("%+v\n", resp.Consignment)
	rsp.Created = true
	rsp.Consignment = consignment
	return nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, rsp *pb.Response) error {
	allConsignments := s.repo.GetAll()
	// resp = &pb.Response{Consignments: allConsignments}	// error
	rsp.Created = true
	rsp.Consignment = allConsignments[0]
	rsp.Consignments = allConsignments

	return nil
}

func main() {
	server := micro.NewService(
		// 必须和 consignment.proto 中的 package 一致
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	// 解析命令行参数
	server.Init()
	repo := Repository{}
	pb.RegisterShippingServiceHandler(server.Server(), &service{repo})

	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
