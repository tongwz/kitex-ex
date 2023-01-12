package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	hello2 "kitex-ex/kitex_gen/hello2/hello2service"
	"kitex-ex/server/hello2/api"
	"kitex-ex/server/hello2_middle/middleware"
	"log"
	"net"
)

func main() {
	// 自定义需要 的配置 例如端口 还有传输协议
	var opts []server.Option
	// 服务端传输协议
	opts = append(opts, server.WithMetaHandler(transmeta.ClientHTTP2Handler))
	// 加入服务端口
	opts = append(opts, server.WithServiceAddr(&net.TCPAddr{Port: 2009}))
	// 服务名称
	opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "hello2Service"}))
	// ***** 加入中间件 查看日志 *****
	opts = append(opts, server.WithMiddleware(middleware.PrintRequestResponseWM))

	svr := hello2.NewServer(new(api.Hello2ServiceImpl), opts...)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
