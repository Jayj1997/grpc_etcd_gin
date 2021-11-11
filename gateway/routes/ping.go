/*
 * @Author       : jayj
 * @Date         : 2021-11-05 14:13:39
 * @Description  :
 */
package routes

import (
	"gateway/proto/pb"
	"gateway/utils/res"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
)

func addPingRoute(rg *gin.RouterGroup) {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBalancerName(roundrobin.Name),
	}

	// 通过etcd做服务发现
	pingConn, err := grpc.Dial("etcd:///tomato/ping", opts...)
	if err != nil {
		logrus.Errorf("try connect ping service failed")
	}
	// 不要调用pingConn.Close() 因为addRoute方法只是导入了路由与对应的方法
	// 在读取完毕之后会关闭链接 在读取完毕之后的请求会无法链接 因为链接已经关闭
	// 如果你一定要严谨的关闭 就把它移出去到公共的地方注册链接 并创建方法在优雅关闭内结束链接

	pingClient := pb.NewPingClient(pingConn)

	ping := rg.Group("/ping")
	{
		ping.POST("", func(ctx *gin.Context) {
			name := ctx.Param("name")

			req := &pb.Req{
				Name: name,
			}

			resp, err := pingClient.Pong(ctx, req)
			if err != nil {
				logrus.Errorln("/v1/ping err, err: ", err)
				res.InternalError(ctx)
				return
			}

			res.Ok(ctx, res.OK, resp.Msg)
		})
	}
}
