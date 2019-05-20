package main

import (
	"context"
	"log"

	pb "shippy/consignment-service/proto/consignment"
	vesselPb "shippy/vessel-service/proto/vessel"

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
	// consignment-service 作为客户端调用 vessel-service 的函数
	vesselClient vesselPb.VesselService
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, rsp *pb.Response) error {
	// 检查是否有适合的货轮
	vReq := &vesselPb.Specification{
		Capacity:  int32(len(req.Containers)),
		MaxWeight: req.Weight,
	}
	vRsp, err := s.vesselClient.FindAvailable(context.Background(), vReq)
	if err != nil {
		return err
	}

	// 货物被承运
	log.Printf("found vessel: %s\n", vRsp.Vessel.Name)
	req.VesselId = vRsp.Vessel.Id
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
	// 作为 vessel-service 的客户端
	vClient := vesselPb.NewVesselService("go.micro.srv.vessel", server.Client())
	pb.RegisterShippingServiceHandler(server.Server(), &service{repo, vClient})

	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
