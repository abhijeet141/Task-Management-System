package main

import (
	"fmt"
	pb "github/grpc-server/proto/generated"
	"io"
	"log"
	"time"

	"github.com/beego/beego/v2/client/orm"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (t *TaskManagementServer) CreateTaskList(stream pb.TaskManagementService_CreateTaskListServer) error {
	o := orm.NewOrm()
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return fmt.Errorf("error receiving task: %v", err)
		}
		createdAt := req.CreatedAt
		if createdAt == nil {
			createdAt = timestamppb.New(time.Now())
		}
		istLocation, err := time.LoadLocation("Asia/Kolkata")
		if err != nil {
			return fmt.Errorf("error loading location %v", err)
		}

		createdAtTimeIst := createdAt.AsTime().In(istLocation).Format("2006-01-02 15:04:05")
		parsedTime, err := time.Parse("2006-01-02 15:04:05", createdAtTimeIst)

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
		log.Printf("Got request with task title: %v", req.Title)

		err = stream.Send(&pb.Task{
			Title:       req.Title,
			Description: req.Description,
			Status:      req.Status,
			CreatedAt:   req.CreatedAt,
		})
		if err != nil {
			return fmt.Errorf("failed to send task %w", err)
		}
		time.Sleep(1 * time.Second)
	}
}
