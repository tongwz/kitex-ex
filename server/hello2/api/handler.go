package api

import (
	"fmt"
	"io"
	hello2 "kitex-ex/kitex_gen/hello2"
)

// Hello2ServiceImpl implements the last service interface defined in the IDL.
type Hello2ServiceImpl struct{}

func (s *Hello2ServiceImpl) ClientSideStreaming(stream hello2.Hello2Service_ClientSideStreamingServer) (err error) {
	println("ClientSideStreaming called")
	// 相当于需要一致 接收客户端发回来的消息 如果需要回复消息 就写回复的消息
	for {
		reqInfo, err := stream.Recv()
		// 表示客户端 关闭了stream消息
		if err == io.EOF {
			continue
		}
		if err != nil {
			fmt.Printf("客户端发送异常： %s \n", err.Error())
			return nil
		}
		fmt.Printf("执行到这儿了：接收到消息是：%#v \n", reqInfo.GetName())
		_ = stream.SendAndClose(&hello2.Response{
			Code: 0,
			Msg:  "周杰伦返回信息~~~" + reqInfo.Name,
		})
	}
}

func (s *Hello2ServiceImpl) ServerSideStreaming(req *hello2.Request, stream hello2.Hello2Service_ServerSideStreamingServer) (err error) {
	println("ServerSideStreaming called")
	fmt.Printf("ServerSideStreaming called info %#v \n", req)
	return stream.Send(&hello2.Response{
		Code: 0,
		Msg:  "童伟珍到此一游~~~" + req.Name,
	})
}

func (s *Hello2ServiceImpl) BidiSideStreaming(stream hello2.Hello2Service_BidiSideStreamingServer) (err error) {
	println("BidiSideStreaming called")
	return
}
