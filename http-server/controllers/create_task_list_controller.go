package controllers

import (
	"context"
	"encoding/json"
	grpcclient "github/http-server/grpc-client"
	pb "github/http-server/proto/generated"
	"io"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

func CreateTaskListController(w http.ResponseWriter, r *http.Request) {
	var tasks pb.TaskList
	TaskId := 0
	err := json.NewDecoder(r.Body).Decode(&tasks)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}
	client, err := grpcclient.TaskManagementClient()
	if err != nil {
		http.Error(w, "Failed to connect to gRPC service", http.StatusInternalServerError)
		return
	}
	stream, err := client.CreateTaskList(context.Background())
	if err != nil {
		http.Error(w, "Failed to create gRPC stream", http.StatusInternalServerError)
		return
	}

	tasksItem := make(map[string]interface{})
	var wg sync.WaitGroup
	var m sync.Mutex

	istLocation, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		http.Error(w, "error loading location %v", http.StatusInternalServerError)
		return
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("error while streaming: %v", err)
				return
			}
			createdTime := res.CreatedAt.AsTime().In(istLocation).Format("2006-01-02 15:04:05")
			parsedTime, err := time.Parse("2006-01-02 15:04:05", createdTime)
			if err != nil {
				http.Error(w, "error parsing formatted time", http.StatusInternalServerError)
				return
			}
			m.Lock()
			tasksItem[strconv.Itoa(TaskId)] = map[string]interface{}{
				"title":       res.Title,
				"description": res.Description,
				"status":      res.Status,
				"created_at":  parsedTime,
				"user_id":     res.UserId,
			}
			TaskId++
			m.Unlock()
		}
	}()
	for _, task := range tasks.Task {
		err := stream.Send(&pb.Task{
			UserId:      task.UserId,
			Title:       task.Title,
			Description: task.Description,
			Status:      task.Status,
		})
		if err != nil {
			log.Fatalf("Error while sending task details: %v", err)
			continue
		}
		log.Printf("Sent Task: %v", task)
		time.Sleep(1 * time.Second)
	}
	stream.CloseSend()
	wg.Wait()
	m.Lock()
	defer m.Unlock()
	response, err := json.MarshalIndent(tasksItem, " ", "\t")
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}
