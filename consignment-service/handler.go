package main

import (
	"context"
	"log"

	pb "shippy/consignment-service/proto/consignment"
	vesselPb "shippy/vessel-service/proto/vessel"

	"gopkg.in/mgo.v2"
)

type handler struct {
	session      *mgo.Session
	vesselClient vesselPb.VesselService
}

// GetRepo 从主会话中 Clone() 出新会话处理查询
func (h *handler) GetRepo() Repository {
	// 	为了避免请求的阻塞，mgo 库提供了 Copy() 和 Clone() 函数来创建新会话，
	// 二者在功能上相差无几，但在细微之处却有重要的区别。Clone 出来的新会话重用了
	// 主会话的 socket，避免了创建 socket 在三次握手时间、资源上的开销，尤其适合
	// 那些快速写入的请求。如果进行了复杂查询、大数据量操作时依旧会阻塞 socket
	// 导致后边的请求阻塞。Copy 为会话创建新的 socket，开销大。
	// 应当根据应用场景不同来选择二者，本文的查询既不复杂数据量也不大，就直接复用
	// 主会话的 socket 即可。不过用完都要 Close()，谨记
	return &ConsignmentRepository{h.session.Clone()}
}

// CreateConsignment ...
func (h *handler) CreateConsignment(ctx context.Context, req *pb.Consignment, rsp *pb.Response) error {
	defer h.GetRepo().Close()

	// 检查是否有适合的货轮
	vReq := &vesselPb.Specification{
		Capacity:  int32(len(req.Containers)),
		MaxWeight: req.Weight,
	}
	vRsp, err := h.vesselClient.FindAvailable(context.Background(), vReq)
	if err != nil {
		return err
	}

	log.Printf("found vessel: %s\n", vRsp.Vessel.Name)
	req.VesselId = vRsp.Vessel.Id
	err = h.GetRepo().Create(req)
	if err != nil {
		return err
	}
	rsp.Created = true
	rsp.Consignment = req
	return nil
}

// GetConsignments ...
func (h *handler) GetConsignments(ctx context.Context, req *pb.GetRequest, rsp *pb.Response) error {
	defer h.GetRepo().Close()
	allConsignments, err := h.GetRepo().GetAll()
	if err != nil {
		return err
	}
	rsp.Created = true
	rsp.Consignment = allConsignments[0]
	rsp.Consignments = allConsignments
	return nil
}
