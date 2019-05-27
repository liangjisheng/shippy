package main

import (
	"context"
	pb "shippy/user-service/proto/user"
)

type handler struct {
	repo Repository
}

// Create ...
func (h *handler) Create(ctx context.Context, req *pb.User, rsp *pb.Response) error {
	if err := h.repo.Create(req); err != nil {
		return err
	}
	rsp.User = req
	return nil
}

// Get ...
func (h *handler) Get(ctx context.Context, req *pb.User, rsp *pb.Response) error {
	u, err := h.repo.Get(req.Id)
	if err != nil {
		return err
	}
	rsp.User = u
	return nil
}

// GetAll ...
func (h *handler) GetAll(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	users, err := h.repo.GetAll()
	if err != nil {
		return err
	}
	rsp.Users = users
	return nil
}

// Auth ...
func (h *handler) Auth(ctx context.Context, req *pb.User, rsp *pb.Token) error {
	_, err := h.repo.GetByEmailAndPassword(req)
	if err != nil {
		return err
	}
	rsp.Token = "x_2nam"
	return nil
}

// ValidateToken ...
func (h *handler) ValidateToken(ctx context.Context, req *pb.Token, rsp *pb.Token) error {
	return nil
}
