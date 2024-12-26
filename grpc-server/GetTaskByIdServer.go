package main

import (
	"context"
	"fmt"
	pb "github/grpc-server/proto/generated"

	"github.com/beego/beego/v2/client/orm"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (t *TaskManagementServer) GetTaskById(ctx context.Context, req *pb.TaskId) (*pb.Task, error) {
	o := orm.NewOrm()
	var task Task
	id := int(req.Id)
	err := o.QueryTable("task").Filter("id", id).One(&task)
	if err != nil {
		return nil, fmt.Errorf("failed to get task %w", err)
	}
	if err == orm.ErrNoRows {
		return nil, fmt.Errorf("no task found with given ID")
	}
	return &pb.Task{
		Id:          uint64(task.Id),
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		CreatedAt:   timestamppb.New(task.CreatedAt)}, nil
}
