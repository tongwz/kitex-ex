// Code generated by Kitex v0.4.4. DO NOT EDIT.

package hello2service

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	transport "github.com/cloudwego/kitex/transport"
	hello2 "kitex-ex/kitex_gen/hello2"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	ClientSideStreaming(ctx context.Context, callOptions ...callopt.Option) (stream Hello2Service_ClientSideStreamingClient, err error)
	ServerSideStreaming(ctx context.Context, Req *hello2.Request, callOptions ...callopt.Option) (stream Hello2Service_ServerSideStreamingClient, err error)
	BidiSideStreaming(ctx context.Context, callOptions ...callopt.Option) (stream Hello2Service_BidiSideStreamingClient, err error)
}

type Hello2Service_ClientSideStreamingClient interface {
	streaming.Stream
	Send(*hello2.Request) error
	CloseAndRecv() (*hello2.Response, error)
}

type Hello2Service_ServerSideStreamingClient interface {
	streaming.Stream
	Recv() (*hello2.Response, error)
}

type Hello2Service_BidiSideStreamingClient interface {
	streaming.Stream
	Send(*hello2.Request) error
	Recv() (*hello2.Response, error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, client.WithTransportProtocol(transport.GRPC))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kHello2ServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kHello2ServiceClient struct {
	*kClient
}

func (p *kHello2ServiceClient) ClientSideStreaming(ctx context.Context, callOptions ...callopt.Option) (stream Hello2Service_ClientSideStreamingClient, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ClientSideStreaming(ctx)
}

func (p *kHello2ServiceClient) ServerSideStreaming(ctx context.Context, Req *hello2.Request, callOptions ...callopt.Option) (stream Hello2Service_ServerSideStreamingClient, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ServerSideStreaming(ctx, Req)
}

func (p *kHello2ServiceClient) BidiSideStreaming(ctx context.Context, callOptions ...callopt.Option) (stream Hello2Service_BidiSideStreamingClient, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.BidiSideStreaming(ctx)
}
