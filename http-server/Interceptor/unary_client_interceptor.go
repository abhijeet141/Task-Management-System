package Interceptor

import (
	"context"
	"log"
	"os"

	"google.golang.org/grpc"
)

var InfoLog *log.Logger
var ErrorLog *log.Logger

func init() {
	file, err := os.OpenFile("Logging/debug.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatalf("Error opening file:%v", err)
	}
	InfoLog = log.New(file, "[INFO] ", log.LstdFlags)
	ErrorLog = log.New(file, "[ERROR] ", log.LstdFlags)
}

func UnaryClientInterceptor(ctx context.Context,
	method string,
	req, reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption) error {

	InfoLog.Printf("Calling method: %s with request: %v", method, req)

	err := invoker(ctx, method, req, reply, cc, opts...)
	if err != nil {
		ErrorLog.Printf("Method %s failed with error: %v", method, err)
	} else {
		InfoLog.Printf("Method %s succeeded with response: %v", method, reply)
	}
	return err
}
