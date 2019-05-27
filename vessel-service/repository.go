package main

import (
	pb "shippy/vessel-service/proto/vessel"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	dbName         = "vessels"
	collectionName = "vessels"
)

// Repository ...
type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
	Create(*pb.Vessel) error
	Close()
}

// VesselRepository ...
type VesselRepository struct {
	session *mgo.Session
}

// FindAvailable ...
func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	// 选择最近一条容量、载重都符合的货轮
	var v *pb.Vessel
	err := repo.collection().Find(bson.M{
		"capacity":  bson.M{"$gte": spec.Capacity},
		"maxweight": bson.M{"$gte": spec.MaxWeight},
	}).One(&v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Create ...
func (repo *VesselRepository) Create(v *pb.Vessel) error {
	return repo.collection().Insert(v)
}

// Close ...
func (repo *VesselRepository) Close() {
	repo.session.Close()
}

func (repo *VesselRepository) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(collectionName)
}
