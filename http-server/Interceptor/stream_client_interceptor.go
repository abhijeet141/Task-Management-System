package Interceptor

import (
	"context"

	"google.golang.org/grpc"
)

type Wrapper struct {
	grpc.ClientStream
	method string
}

func (w *Wrapper) RecvMsg(m interface{}) error {
	err := w.ClientStream.RecvMsg(m)
	if err == nil {
		InfoLog.Printf("Received message from method %s: %v", w.method, m)
	} else {
		ErrorLog.Printf("Error receiving message from method %s: %v", w.method, err)
	}
	return err
}
func StreamClientInterceptor(ctx context.Context,
	desc *grpc.StreamDesc,
	cc *grpc.ClientConn,
	method string,
	streamer grpc.Streamer,
	opts ...grpc.CallOption,
) (grpc.ClientStream, error) {
	InfoLog.Printf("Stream successfully opened for method: %s", method)
	clientStream, err := streamer(ctx, desc, cc, method, opts...)
	if err != nil {
		ErrorLog.Printf("Stream creation failed for method %s: %v", method, err)
		return nil, err
	}
	InfoLog.Printf("Stream successfully closed for method: %s", method)
	return &Wrapper{ClientStream: clientStream, method: method}, err
}
