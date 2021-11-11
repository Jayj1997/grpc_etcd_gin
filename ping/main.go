/*
 * @Author       : jayj
 * @Date         : 2021-11-04 14:41:10
 * @Description  :
 */
package main

import (
	"fmt"
	"net"
	"ping/discovery"
	"ping/proto/pb"
	"ping/service"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	app         = "tomato/ping"
	grpcAddress = "127.0.0.1:10004"
)

func main() {
	etcdAddress := []string{"127.0.0.1:2379"}

	etcdRegister := discovery.NewRegister(etcdAddress, logrus.New())
	defer etcdRegister.Stop()

	node := discovery.Server{
		Name: app,
		Addr: grpcAddress,
	}

	server := grpc.NewServer()
	defer server.Stop()

	pb.RegisterPingServer(server, service.NewPingService())

	lis, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		panic(err)
	}
	defer lis.Close()

	if _, err := etcdRegister.Register(node, 10); err != nil {
		panic(fmt.Sprintf("start server failed, err: %v", err))
	}

	logrus.Info("server started listen on ", grpcAddress)
	if err := server.Serve(lis); err != nil {
		panic(err)
	}
}
