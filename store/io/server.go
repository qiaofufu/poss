package io

import (
	"OSS/store/proto"
	"context"
)

type storeServer struct {
}

func Run() {

}

func (s *storeServer) Mkdir(ctx context.Context, param *proto.MkdirParam) (*proto.MkdirResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *storeServer) Rmdir(ctx context.Context, param *proto.RmdirParam) (*proto.RmdirResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *storeServer) Stat(ctx context.Context, param *proto.StatParam) (*proto.StatResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *storeServer) Open(ctx context.Context, param *proto.OpenParam) (*proto.OpenResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *storeServer) Write(server proto.Storage_WriteServer) error {
	//TODO implement me
	panic("implement me")
}

func (s *storeServer) Read(param *proto.ReadParam, server proto.Storage_ReadServer) error {
	//TODO implement me
	panic("implement me")
}

func (s *storeServer) Del(ctx context.Context, param *proto.DelParam) (*proto.DelResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *storeServer) mustEmbedUnimplementedStorageServer() {
	//TODO implement me
	panic("implement me")
}
