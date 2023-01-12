package main

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/remote/codec/thrift"
	"github.com/cloudwego/kitex/pkg/transmeta"
	hello2 "kitex-ex/kitex_gen/hello"
	hello "kitex-ex/kitex_gen/hello/helloservice"
	"log"
	"strconv"
	"time"
)

func main() {
	clientCustomer, err := hello.NewClient(
		"tongWzHello",
		client.WithHostPorts("0.0.0.0:2008"),
		// 解析协议
		client.WithPayloadCodec(thrift.NewThriftCodecWithConfig(thrift.FastRead|thrift.FastWrite)),
		// 多路复用
		client.WithMuxConnection(2),
		// 传输协议 使用TTHeader
		client.WithMetaHandler(transmeta.ClientTTHeaderHandler),
	)
	if err != nil {
		log.Fatal(err)
	}
	i := 0
	for {
		req := &hello2.Request{
			Message: "t request " + strconv.Itoa(i),
			Data:    "t request data" + strconv.Itoa(i),
			ReqBody: &hello2.ReqBody{
				Name: "hello" + strconv.Itoa(i),
				Type: 1,
			},
		}
		resp, err := clientCustomer.Echo(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp)
		time.Sleep(time.Millisecond * 1)
		i++
	}
}
