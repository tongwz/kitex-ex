package main

import (
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/remote/codec"
	"github.com/cloudwego/kitex/pkg/remote/codec/thrift"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	hello "kitex-ex/kitex_gen/hello/helloservice"
	"kitex-ex/server/hello/api"
	"log"
	"net"
)

func main() {
	svr := hello.NewServer(new(api.HelloServiceImpl),
		// 设置服务端口
		server.WithServiceAddr(&net.TCPAddr{Port: 2008}),

		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "tongWzHello"}),
		server.WithPayloadCodec(thrift.NewThriftCodecWithConfig(thrift.FastRead|thrift.FastWrite)),

		// 指定默认 Codec 的包大小限制，默认无限制 option: codec.NewDefaultCodecWithSizeLimit
		server.WithCodec(codec.NewDefaultCodecWithSizeLimit(1024*1024*10)), //10M
		server.WithLimit(&limit.Option{MaxConnections: 10000, MaxQPS: 5000}),
		// 连接多路复用(mux)
		server.WithMuxTransport(),
		server.WithMetaHandler(transmeta.ServerTTHeaderHandler),
	)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
