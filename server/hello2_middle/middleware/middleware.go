package middleware

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/endpoint"
)

func PrintRequestResponseWM(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		fmt.Printf("Request : %v \n", req)
		err = next(ctx, req, resp)
		fmt.Printf("Response : %v \n", resp)
		return err
	}
}
