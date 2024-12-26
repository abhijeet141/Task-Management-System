package main

import (
	"context"
	"fmt"
	pb "github/grpc-server/proto/generated"

	"github.com/beego/beego/v2/client/orm"
)

func (t *TaskManagementServer) DeleteTaskById(ctx context.Context, req *pb.TaskId) (*pb.Message, error) {
	o := orm.NewOrm()
	id := int(req.Id)
	num, err := o.QueryTable("task").Filter("id", id).Delete()
	if err != nil {
		return nil, fmt.Errorf("failed to delete task %w", err)
	}
	if num == 0 {
		return nil, fmt.Errorf("Task not found %w", err)
	}
	return &pb.Message{Message: "Task Deleted Successfully"}, nil
}
