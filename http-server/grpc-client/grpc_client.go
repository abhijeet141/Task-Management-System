package grpcclient

import (
	"github/http-server/Interceptor"
	pb "github/http-server/proto/generated"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	PORT = ":8080"
)

func TaskManagementClient() (pb.TaskManagementServiceClient, error) {
	conn, err := grpc.NewClient("localhost"+PORT, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithUnaryInterceptor(Interceptor.UnaryClientInterceptor), grpc.WithStreamInterceptor(Interceptor.StreamClientInterceptor))
	if err != nil {
		log.Fatalf("Did not connect %v", err)
	}
	client := pb.NewTaskManagementServiceClient(conn)
	return client, nil
}
