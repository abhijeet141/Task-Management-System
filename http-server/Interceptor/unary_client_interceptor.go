package interceptor

import (
	"context"
	"fmt"
	"log"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func UnaryClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	token := r.Header.Get("Authorization")
	if token == "" || !strings.HasPrefix(token, "Bearer ") {
		return fmt.Errorf("Invalid Authorization Header")
	}
	tokenId := strings.TrimPrefix(token, "Bearer ")
	ctx = metadata.AppendToOutgoingContext(ctx, "tokenId", tokenId)
	err := invoker(ctx, method, req, reply, cc, opts...)
	if err != nil {
		log.Printf("failed: %v", err)
	} else {
		log.Printf("Response received: %v", reply)
	}
	return err
}
