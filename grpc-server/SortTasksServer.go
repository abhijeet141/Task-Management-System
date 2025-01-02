package main

import (
	"context"
	"fmt"
	"github/grpc-server/Reflection"
	pb "github/grpc-server/proto/generated"

	"github.com/beego/beego/v2/client/orm"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (t *TaskManagementServer) SortTasks(ctx context.Context, req *pb.SortTasksRequest) (*pb.TaskList, error) {
	o := orm.NewOrm()
	sortBy := req.SortBy
	if sortBy == "" {
		sortBy = "Id"
	}
	var tasks []Task
	_, err := o.QueryTable("task").All(&tasks)
	if err != nil {
		return nil, err
	}
	err = Reflection.DynamicSort(&tasks, sortBy)
	if err != nil {
		return nil, fmt.Errorf("error in sorting")
	}
	var pbTasks []*pb.Task
	for _, task := range tasks {
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
