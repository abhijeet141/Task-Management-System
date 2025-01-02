package controllers

import (
	"context"
	"encoding/json"
	grpcclient "github/http-server/grpc-client"
	pb "github/http-server/proto/generated"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Task struct {
	Id          int       `orm:"column(id);auto"`
	UserId      int       `orm:"column(user_id)"`
	Title       string    `orm:"column(title)"`
	Description string    `orm:"column(description)"`
	Status      string    `orm:"column(status)"`
	CreatedAt   time.Time `orm:"column(created_at);type(datetime)"`
}

func SortTasksControllers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sortBy := vars["sortBy"]
	if sortBy != "Status" && sortBy != "CreatedAt" && sortBy != "Title" && sortBy != "Id" {
		http.Error(w, "Invalid sort criteria", http.StatusBadRequest)
		return
	}
	sortTaskReq := &pb.SortTasksRequest{SortBy: sortBy}
	client, err := grpcclient.TaskManagementClient()
	if err != nil {
		http.Error(w, "Failed to connect to gRPC service", http.StatusInternalServerError)
		return
	}
	res, err := client.SortTasks(context.Background(), sortTaskReq)
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}
	istLocation, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		http.Error(w, "Error loading location %v", http.StatusInternalServerError)
		return
	}
	var Tasks []Task
	for _, task := range res.Task {
		createdTime := task.CreatedAt.AsTime().In(istLocation).Format("2006-01-02 15:04:05")
		parsedTime, err := time.Parse("2006-01-02 15:04:05", createdTime)
		if err != nil {
			http.Error(w, "error parsing formatted time", http.StatusInternalServerError)
			return
		}
		Task := Task{
			Id:          int(task.Id),
			Title:       task.Title,
			Description: task.Description,
			Status:      task.Status,
			CreatedAt:   parsedTime,
		}
		Tasks = append(Tasks, Task)
	}
	response, err := json.MarshalIndent(Tasks, "", "\t")
	if err != nil {
		log.Fatalf("Error marshalling JSON: %v", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}
