/*
 * @Author       : jayj
 * @Date         : 2021-11-11 15:18:13
 * @Description  :
 */
package main

import (
	"gateway/discovery"
	"gateway/routes"
	"gateway/utils"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/resolver"
)

func main() {

	// etcd注册
	r := discovery.NewResolver([]string{"localhost:2379"}, logrus.New())
	defer r.Close()
	resolver.Register(r)

	g := routes.LoadGin()

	server := &http.Server{
		Addr:           ":8080",
		Handler:        g,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	logrus.Info("gateway listen on :8080")

	go func() {
		// 优雅关闭
		utils.GracefullyShutdown(server)
	}()

	if err := server.ListenAndServe(); err != nil {
		logrus.Fatal("gateway启动失败, err: ", err)
	}
}
