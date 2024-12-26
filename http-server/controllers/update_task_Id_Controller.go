package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	grpcclient "github/http-server/grpc-client"
	pb "github/http-server/proto/generated"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func UpdateTaskByIdController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}
	taskId := &pb.TaskId{Id: id}
	var task *pb.Task
	err = json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}
	taskItem := &pb.Task{
		Id:          taskId.Id,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
	}
	client, err := grpcclient.TaskManagementClient()
	if err != nil {
		http.Error(w, "Failed to connect to gRPC service", http.StatusInternalServerError)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.UpdateTaskById(ctx, taskItem)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to update task: %v", err), http.StatusInternalServerError)
		return
	}
	istLocation, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		http.Error(w, "error loading location %v", http.StatusInternalServerError)
		return
	}
	createdTime := res.CreatedAt.AsTime().In(istLocation).Format("2006-01-02 15:04:05")
	parsedTime, err := time.Parse("2006-01-02 15:04:05", createdTime)
	if err != nil {
		http.Error(w, "error parsing formatted time", http.StatusInternalServerError)
		return
	}
	response, err := json.MarshalIndent(map[string]map[string]interface{}{
		"Server Response": {
			"Id":          res.Id,
			"Title":       res.Title,
			"Description": res.Description,
			"Status":      res.Status,
			"Created At":  parsedTime,
		},
	}, "", "\t")
	if err != nil {
		log.Fatalf("Error marshalling JSON: %v", err)
	}
	log.Println(string(response))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}
