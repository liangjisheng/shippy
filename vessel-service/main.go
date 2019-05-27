package main

import (
	"log"
	"os"
	pb "shippy/vessel-service/proto/vessel"

	"github.com/micro/go-micro"
)

const (
	defaultDBHost = "localhost:27017"
)

func main() {
	// 获取容器设置的数据库地址环境变量的值
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = defaultDBHost
	}
	log.Println("dbHost:", dbHost)
	session, err := CreateSession(dbHost)
	if err != nil {
		log.Fatalf("create session error:%v\n", err)
	}
	defer session.Close()

	// 停留在港口的货船，先写死
	repo := &VesselRepository{session.Copy()}
	CreateDummyData(repo)

	server := micro.NewService(
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
	)
	server.Init()

	pb.RegisterVesselServiceHandler(server.Server(), &handler{session})

	if err := server.Run(); err != nil {
		log.Fatalf("failed to server: %v\n", err)
	}
}

// CreateDummyData ...
func CreateDummyData(repo Repository) {
	defer repo.Close()
	vessels := []*pb.Vessel{
		{Id: "vessel001", Name: "Boaty McBoatface", MaxWeight: 200000, Capacity: 500},
	}
	for _, v := range vessels {
		repo.Create(v)
	}
}
