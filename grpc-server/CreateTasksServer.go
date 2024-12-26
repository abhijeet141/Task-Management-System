package main

import (
	"fmt"
	pb "github/grpc-server/proto/generated"
	"io"
	"time"

	"github.com/beego/beego/v2/client/orm"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (t *TaskManagementServer) CreateTasks(stream pb.TaskManagementService_CreateTasksServer) error {
	for {
		o := orm.NewOrm()
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.Message{Message: "Tasks Created Successfully"})
		}

		if err != nil {
			return fmt.Errorf("error receiving task: %v", err)
		}

		createdAt := req.CreatedAt
		if createdAt == nil {
			createdAt = timestamppb.New(time.Now())
		}

		createdAtTime := createdAt.AsTime()
		istLocation, err := time.LoadLocation("Asia/Kolkata")
		if err != nil {
			return fmt.Errorf("error loading location %v", err)
		}

		createdAtTimeIst := createdAtTime.In(istLocation)
		formattedTime := createdAtTimeIst.Format("2006-01-02 15:04:05")

		parsedTime, err := time.Parse("2006-01-02 15:04:05", formattedTime)
		if err != nil {
			return fmt.Errorf("error parsing formatted time: %v", err)
		}

		dbTask := Task{
			Title:       req.Title,
			Description: req.Description,
			Status:      req.Status,
			CreatedAt:   parsedTime,
			UserId:      int(req.UserId),
		}

		_, err = o.Insert(&dbTask)
		if err != nil {
			return fmt.Errorf("failed to insert task %w", err)
		}
	}
}
