package service

import (
	"context"
	"fmt"
	"ping/proto/pb"
)

type PingService struct {
}

func NewPingService() *PingService {
	return &PingService{}
}

func (s *PingService) Pong(ctx context.Context, in *pb.Req) (*pb.Resp, error) {

	msg := fmt.Sprintf("hello %s", in.Name)

	fmt.Println("received")

	return &pb.Resp{
		Msg: msg,
	}, nil
}
