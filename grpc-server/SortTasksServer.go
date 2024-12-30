package main

import (
	"context"
	"fmt"
	pb "github/grpc-server/proto/generated"

	"github.com/beego/beego/v2/client/orm"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (t *TaskManagementServer) SortTasks(ctx context.Context, req *pb.SortTasksRequest) (*pb.TaskList, error) {
	o := orm.NewOrm()
	sortBy := req.SortBy
	if sortBy == "" {
		sortBy = "created_at"
	}
	var tasks []Task
	_, err := o.QueryTable("task").OrderBy(sortBy).All(&tasks)
	if err != nil {
		return nil, err
	}
	fmt.Println(tasks)
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
