package api

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	hello "kitex-ex/kitex_gen/hello"
	"net"
	"strings"
)

// HelloServiceImpl implements the last service interface defined in the IDL.
type HelloServiceImpl struct{}

// Echo implements the HelloServiceImpl interface.
func (s *HelloServiceImpl) Echo(ctx context.Context, req *hello.Request) (resp *hello.Response, err error) {
	// 具体服务端逻辑
	klog.Info("hello service enter:" + GetIpAddr2())
	resp = &hello.Response{
		Msg: &hello.Msg{
			Status: 200,
			Code:   0,
			Msg:    req.Message,
		},
		Data: req.Message,
	}
	return resp, nil
}

func GetIpAddr2() string {
	// tongWz??? 这个地址为啥这么写不知道
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		klog.Error(err)
		return ""
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	// 192.168.1.20:61085
	ip := strings.Split(localAddr.String(), ":")[0]

	return ip
}

// TestHello4Get implements the HelloServiceImpl interface.
func (s *HelloServiceImpl) TestHello4Get(ctx context.Context, req *hello.Request) (resp *hello.Response, err error) {
	// TODO: Your code here...
	return
}

// TestHello4Post implements the HelloServiceImpl interface.
func (s *HelloServiceImpl) TestHello4Post(ctx context.Context, req *hello.Request) (resp *hello.Response, err error) {
	// TODO: Your code here...
	return
}
