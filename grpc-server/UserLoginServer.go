package main

import (
	"context"
	"fmt"
	pb "github/grpc-server/proto/generated"

	"github.com/beego/beego/v2/client/orm"
)

func (t *TaskManagementServer) UserLogin(ctx context.Context, req *pb.UserInfo) (*pb.Message, error) {
	o := orm.NewOrm()
	userName := User{EmailAddress: req.UserName}
	err := o.Read(&userName, "email_address")
	if err != nil {
		return nil, fmt.Errorf("user not found %v", err)
	}
	if userName.Password != req.Password {
		return nil, fmt.Errorf("invalid password")
	}
	return &pb.Message{
		Message: "Login successful",
	}, nil
}
