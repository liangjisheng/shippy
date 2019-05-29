package main

import (
	"context"
	"errors"
	"log"
	pb "shippy/user-service/proto/user"

	"github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/broker/nats"
	"golang.org/x/crypto/bcrypt"
)

const topic = "user.created"

type handler struct {
	repo         Repository
	tokenService Authable
	Publisher    micro.Publisher
	// PubSub broker.Broker
}

// Create ...
func (h *handler) Create(ctx context.Context, req *pb.User, rsp *pb.Response) error {
	// 哈希处理用户输入的密码
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req.Password = string(hashedPwd)

	if err := h.repo.Create(req); err != nil {
		return err
	}
	rsp.User = req

	// 发布带有用户所有信息的消息
	// if err := h.publishEvent(req); err != nil {
	// 	return err
	// }

	if err := h.Publisher.Publish(ctx, req); err != nil {
		return err
	}
	return nil
}

// publishEvent 发送消息通知
// func (h *handler) publishEvent(user *pb.User) error {
// 	body, err := json.Marshal(user)
// 	if err != nil {
// 		return err
// 	}

// 	msg := &broker.Message{
// 		Header: map[string]string{
// 			"id": user.Id,
// 		},
// 		Body: body,
// 	}

// 	// 发布 user.created topic 消息
// 	if err := h.PubSub.Publish(topic, msg); err != nil {
// 		log.Fatalf("[pub] failed: %v\n", err)
// 	}
// 	return nil
// }

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
	// 在 part3 中直接传参 &pb.User 去查找用户
	// 会导致 req 的值完全是数据库中的记录值
	// 即 req.Password 与 u.Password 都是加密后的密码
	// 将无法通过验证
	u, err := h.repo.GetByEmail(req.Email)
	if err != nil {
		log.Println("GetByEmailAndPassword", err)
		return err
	}

	// 进行密码验证
	log.Println("u.password:", u.Password)
	log.Println("req.password:", req.Password)
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password))
	if err != nil {
		log.Println("CompareHashAndPassword", err)
		return err
	}

	t, err := h.tokenService.Encode(u)
	if err != nil {
		log.Println("tokenService.Encode", err)
		return err
	}
	rsp.Token = t
	return nil
}

// ValidateToken ...
func (h *handler) ValidateToken(ctx context.Context, req *pb.Token, rsp *pb.Token) error {
	// Decode token
	claims, err := h.tokenService.Decode(req.Token)
	if err != nil {
		return err
	}
	if "" == claims.User.Id {
		return errors.New("invalid user")
	}

	rsp.Valid = true
	rsp.Token = req.Token
	return nil
}
