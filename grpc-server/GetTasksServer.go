package main

import (
	"context"
	"fmt"
	pb "github/grpc-server/proto/generated"

	"github.com/beego/beego/v2/client/orm"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (t *TaskManagementServer) GetAllTask(ctx context.Context, req *pb.NoParam) (*pb.TaskList, error) {
	o := orm.NewOrm()
	var task []Task
	_, err := o.QueryTable("task").All(&task)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve task %w", err)
	}
	if err == orm.ErrNoRows {
		return nil, fmt.Errorf("no task found with given ID")
	}
	var pbTasks []*pb.Task
	for _, task := range task {
		pbTask := &pb.Task{
			Id:          uint64(task.Id),
			Title:       task.Title,
			Description: task.Description,
			Status:      task.Status,
			CreatedAt:   timestamppb.New(task.CreatedAt),
		}
		pbTasks = append(pbTasks, pbTask)
	}
	return &pb.TaskList{Task: pbTasks}, nil
}
