package main

import (
	"context"
	"errors"
	"log"

	pb "shippy/vessel-service/proto/vessel"

	"github.com/micro/go-micro"
)

// Repository ...
type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
}

// VesselRepository ...
type VesselRepository struct {
	vessels []*pb.Vessel
}

// FindAvailable ...
func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	// 选择最近一条容量、载重都符合的货轮
	for _, v := range repo.vessels {
		if v.Capacity >= spec.Capacity && v.MaxWeight >= spec.MaxWeight {
			return v, nil
		}
	}

	return nil, errors.New("No vessel can't be use")
}

type service struct {
	repo Repository
}

func (s *service) FindAvailable(ctx context.Context, spec *pb.Specification, rsp *pb.Response) error {
	// 调用内部方法查找
	v, err := s.repo.FindAvailable(spec)
	if err != nil {
		return err
	}
	rsp.Vessel = v
	return nil
}

func main() {
	// 停留在港口的货船，先写死
	vessels := []*pb.Vessel{
		{Id: "vessel001", Name: "Boaty McBoatface", MaxWeight: 200000, Capacity: 500},
	}
	repo := &VesselRepository{vessels}

	server := micro.NewService(
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
	)
	server.Init()

	pb.RegisterVesselServiceHandler(server.Server(), &service{repo})

	if err := server.Run(); err != nil {
		log.Fatalf("failed to server: %v\n", err)
	}
}
