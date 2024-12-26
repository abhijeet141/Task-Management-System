package main

import (
	"context"
	"fmt"
	pb "github/grpc-server/proto/generated"
	"time"

	"github.com/beego/beego/v2/client/orm"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (t *TaskManagementServer) UserRegister(ctx context.Context, req *pb.User) (*pb.Message, error) {
	o := orm.NewOrm()
	istLocation, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		return nil, fmt.Errorf("error loading location %v", err)
	}
	createdAt := req.CreatedAt
	if createdAt == nil {
		createdAt = timestamppb.New(time.Now())
	}
	createdAtTime := createdAt.AsTime().In(istLocation).Format("2006-01-02 15:04:05")
	parsedTime, err := time.Parse("2006-01-02 15:04:05", createdAtTime)
	if err != nil {
		return nil, fmt.Errorf("error parsing formatted time: %v", err)
	}
	dbUser := User{
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		EmailAddress: req.EmailAddress,
		Password:     req.Password,
		CreatedAt:    parsedTime,
	}
	_, err = o.Insert(&dbUser)
	if err != nil {
		return nil, fmt.Errorf("failed to insert user %w", err)
	}
	return &pb.Message{
		Message: "User Created Successfully",
	}, nil
}
