package main

import (
	"context"
	"fmt"
	pb "github/grpc-server/proto/generated"
	"time"

	"github.com/beego/beego/v2/client/orm"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (t *TaskManagementServer) CreateTask(ctx context.Context, req *pb.Task) (*pb.Message, error) {
	o := orm.NewOrm()
	createdAt := req.CreatedAt
	if createdAt == nil {
		createdAt = timestamppb.New(time.Now())
	}

	createdAtTime := createdAt.AsTime()

	dbTask := Task{
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
		CreatedAt:   createdAtTime,
	}

	_, err := o.Insert(&dbTask)
	if err != nil {
		return nil, fmt.Errorf("failed to insert task %w", err)
	}
	return &pb.Message{
		Message: "Task Created Successfully",
	}, nil
}
