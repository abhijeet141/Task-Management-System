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
	file, err := os.OpenFile("Logging/debug.log", os.O_APPEND|os.O_CREATE|os.O_RDONLY, 0444)
	if err != nil {
		log.Fatalf("Error opening file:%v", err)
	}
	InfoLog = log.New(file, "[INFO] ", log.LstdFlags)
	ErrorLog = log.New(file, "[ERROR] ", log.LstdFlags)
}
func UnaryServerInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	InfoLog.Printf("Calling method: %s with request: %v", info.FullMethod, req)
	response, err := handler(ctx, req)
	if err != nil {
		ErrorLog.Printf("Method %s failed with error: %v", info.FullMethod, err)
	} else {
		InfoLog.Printf("Method %s succeeded with response: %v", info.FullMethod, response)
	}
	return response, err
}
