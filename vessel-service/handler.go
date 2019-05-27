package main

import (
	"context"
	pb "shippy/vessel-service/proto/vessel"

	"gopkg.in/mgo.v2"
)

type handler struct {
	session *mgo.Session
}

// GetRepo ...
func (h *handler) GetRepo() Repository {
	return &VesselRepository{h.session.Clone()}
}

// FindAvailable ...
func (h *handler) FindAvailable(ctx context.Context, req *pb.Specification, rsp *pb.Response) error {
	defer h.GetRepo().Close()
	v, err := h.GetRepo().FindAvailable(req)
	if err != nil {
		return err
	}
	rsp.Vessel = v
	return nil
}

// Create ...
func (h *handler) Create(ctx context.Context, req *pb.Vessel, rsp *pb.Response) error {
	defer h.GetRepo().Close()
	if err := h.GetRepo().Create(req); err != nil {
		return err
	}
	rsp.Vessel = req
	rsp.Created = true
	return nil
}
