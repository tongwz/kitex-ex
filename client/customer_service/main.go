package main

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/remote/codec/thrift"
	hello2 "kitex-ex/kitex_gen/hello"
	hello "kitex-ex/kitex_gen/hello/helloservice"
	"log"
	"time"
)

func main() {
	clientCustomer, err := hello.NewClient(
		"tongWzHello",
		client.WithHostPorts("0.0.0.0:2008"),
		// 解析协议
		client.WithPayloadCodec(thrift.NewThriftCodecWithConfig(thrift.FastRead|thrift.FastWrite)),
		// 多路复用
		client.WithMuxConnection(3),
		//client.WithMetaHandler(transmeta.ClientTTHeaderHandler),
	)
	if err != nil {
		log.Fatal(err)
	}
	for {
		req := &hello2.Request{
			Message: "t request 1",
			Data:    "t request data1",
			ReqBody: &hello2.ReqBody{
				Name: "hello",
				Type: 1,
			},
		}
		resp, err := clientCustomer.Echo(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp)
		time.Sleep(time.Second)
	}
}
