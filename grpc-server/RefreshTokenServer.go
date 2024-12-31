package main

import (
	"context"
	"fmt"
	pb "github/grpc-server/proto/generated"
	"time"

	"github.com/beego/beego/v2/client/orm"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (t *TaskManagementServer) RefreshToken(ctx context.Context, req *pb.Token) (*pb.Message, error) {
	o := orm.NewOrm()
	createdAt := timestamppb.New(time.Now())
	createdAtTime := createdAt.AsTime()
	istLocation, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		return nil, fmt.Errorf("error loading location %v", err)
	}
	createdAtTimeIst := createdAtTime.In(istLocation)
	formattedTime := createdAtTimeIst.Format("2006-01-02 15:04:05")
	parsedTime, err := time.Parse("2006-01-02 15:04:05", formattedTime)
	if err != nil {
		return nil, fmt.Errorf("error parsing formatted time: %v", err)
	}
	refToken := Token{
		UserId:    int(req.UserId),
		JWT:       req.Jwt,
		CreatedAt: parsedTime,
		Expired:   req.Expired,
	}
	_, err = o.Insert(&refToken)
	if err != nil {
		return nil, fmt.Errorf("failed to insert refresh token %w", err)
	}
	return &pb.Message{
		Message: "Refresh Token inserted Successfully",
	}, nil
}
