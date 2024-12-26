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
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		return fmt.Errorf("no metadata in context")
	}
	token := ""
	if values, exists := md["Authorization"]; exists && len(values) > 0 {
		token = values[0]
	}
	if token == "" || !strings.HasPrefix(token, "Bearer ") {
		return fmt.Errorf("invalid Authorization Header")
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
