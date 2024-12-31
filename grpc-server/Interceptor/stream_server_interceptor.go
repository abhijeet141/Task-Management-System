package Interceptor

import (
	"google.golang.org/grpc"
)

type wrappedServerStream struct {
	grpc.ServerStream
}

func (w *wrappedServerStream) RecvMsg(m interface{}) error {
	err := w.ServerStream.RecvMsg(m)
	if err != nil {
		ErrorLog.Printf("Error receiving message: %v", err)
	} else {
		InfoLog.Printf("Received message: %v", m)
	}
	return err
}
func StreamServerInterceptor(srv interface{},
	ss grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {
	InfoLog.Printf("Stream sucessfully opened for method: %s", info.FullMethod)
	err := handler(srv, &wrappedServerStream{ServerStream: ss})
	if err != nil {
		ErrorLog.Printf("Stream creation failed for method %s: %v", info.FullMethod, err)
		return err
	}
	InfoLog.Printf("Stream successfully closed for method: %s", info.FullMethod)
	return err
}
