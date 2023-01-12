package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/remote/codec/protobuf"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	hello22 "kitex-ex/kitex_gen/hello2"
	hello2 "kitex-ex/kitex_gen/hello2/hello2service"

	"log"
	"strconv"
	"time"
)

func main() {
	var opts []client.Option
	// 传输协议 使用GRPC
	opts = append(opts, client.WithTransportProtocol(transport.GRPC))
	// 原数据 使用协议 http2
	opts = append(opts, client.WithMetaHandler(transmeta.ClientHTTP2Handler))
	opts = append(opts, client.WithHostPorts("0.0.0.0:2009"))
	// 解析协议
	opts = append(opts, client.WithPayloadCodec(protobuf.NewProtobufCodec()))
	// 多路复用
	opts = append(opts, client.WithMuxConnection(2))
	clientHello2, err := hello2.NewClient("hello2Service", opts...)
	if err != nil {
		log.Fatal(err)
	}
	cxx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		j := 0
		for {
			// 客户端侧 streaming
			streamObj, err := clientHello2.ClientSideStreaming(cxx)
			if err != nil {
				log.Fatal(err)
			}
			err = streamObj.Send(
				&hello22.Request{
					Name: "测试客户端Streaming " + strconv.Itoa(j),
					Id:   int32(j),
					ReqBody: []*hello22.ReqBody{
						&hello22.ReqBody{
							Number: "测试客户端Streaming~" + strconv.Itoa(j),
							Id:     int32(j),
						},
					},
				})
			if err != nil {
				log.Fatal(err)
			}

			respClientSideStreaming, _ := streamObj.CloseAndRecv()
			fmt.Printf("-----------------ClientSideStreaming我们获取到的返回值是 code:%d, msg: %s \n", respClientSideStreaming.Code, respClientSideStreaming.Msg)
			time.Sleep(time.Millisecond * 2000)
			j++
		}
	}()
	i := 0
	for {
		req := &hello22.Request{
			Name: "偶像周杰伦 " + strconv.Itoa(i),
			Id:   int32(i),
			ReqBody: []*hello22.ReqBody{
				&hello22.ReqBody{
					Number: "童伟珍测试~" + strconv.Itoa(i),
					Id:     int32(i),
				},
			},
		}
		resp, err := clientHello2.ServerSideStreaming(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		rev, _ := resp.Recv()
		fmt.Printf("我们获取到的返回值是 code:%d, msg: %s \n", rev.Code, rev.Msg)
		time.Sleep(time.Millisecond * 2000)
		i++
	}
}
