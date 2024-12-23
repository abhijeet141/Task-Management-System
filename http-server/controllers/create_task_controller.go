package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	grpcclient "github/http-server/grpc-client"
	pb "github/http-server/proto/generated"
	"log"
	"net/http"
)

func CreateTaskController(w http.ResponseWriter, r *http.Request) {
	var task pb.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}
	taskItem := &pb.Task{
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
	}

	client, err := grpcclient.TaskManagementClient()
	if err != nil {
		http.Error(w, "Failed to connect to gRPC service", http.StatusInternalServerError)
		return
	}
	res, err := client.CreateTask(context.Background(), taskItem)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create task: %v", err), http.StatusInternalServerError)
		return
	}
	response, err := json.MarshalIndent(map[string]string{"Server Response": res.Message}, "", " ")
	if err != nil {
		log.Fatalf("Error Marshalling Response %v", err)
	}
	log.Println(string(response))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}
