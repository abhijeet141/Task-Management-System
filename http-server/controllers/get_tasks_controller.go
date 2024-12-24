package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	grpcclient "github/http-server/grpc-client"
	pb "github/http-server/proto/generated"
	"log"
	"net/http"
	"time"
)

func GetTasksController(w http.ResponseWriter, r *http.Request) {
	client, err := grpcclient.TaskManagementClient()
	if err != nil {
		http.Error(w, "Failed to connect to gRPC service", http.StatusInternalServerError)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.GetAllTask(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}
	tasks := map[string]interface{}{}
	istLocation, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		http.Error(w, "error loading location %v", http.StatusInternalServerError)
		return
	}

	for _, task := range res.Task {
		createdTime := task.CreatedAt.AsTime().In(istLocation).Format("2006-01-02 15:04:05")
		parsedTime, err := time.Parse("2006-01-02 15:04:05", createdTime)
		if err != nil {
			http.Error(w, "error parsing formatted time", http.StatusInternalServerError)
			return
		}
		tasks[fmt.Sprintf("%d", task.Id)] = map[string]interface{}{
			"title":       task.Title,
			"description": task.Description,
			"status":      task.Status,
			"created_at":  parsedTime,
		}
	}
	response, err := json.MarshalIndent(tasks, "", "\t")
	if err != nil {
		log.Fatalf("Error marshalling JSON: %v", err)
	}
	log.Println(string(response))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}
