package middleware

import (
	"context"
	"github.com/cloudwego/kitex/pkg/endpoint"
)

func GatewayMiddleware(bucket *Bucket) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, req, resp interface{}) (err error) {
			// 检查请求是否在活动时间内
			if !isInActivityTime() {
				return &MyError{Code: 400, Message: "请求不在活动时间内"}
			}
			// 进行限流检查
			if !bucket.Allow() {
				return &MyError{Code: 429, Message: "请求过于频繁，请稍后再试"}
			}
			// 调用下一个中间件或实际的服务方法
			return next(ctx, req, resp)
		}
	}
}

// MyError 自定义错误结构体
type MyError struct {
	Code    int
	Message string
}

func (e *MyError) Error() string {
	return e.Message
}
