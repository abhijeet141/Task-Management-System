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

	"github.com/gorilla/mux"
)

func SortTasksControllers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sortBy := vars["sortBy"]
	if sortBy != "status" && sortBy != "created_at" {
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
		http.Error(w, "error loading location %v", http.StatusInternalServerError)
		return
	}
	tasks := map[string]interface{}{}
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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}