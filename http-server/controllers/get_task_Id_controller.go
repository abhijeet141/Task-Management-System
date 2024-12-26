package controllers

import (
	"context"
	"encoding/json"
	grpcclient "github/http-server/grpc-client"
	pb "github/http-server/proto/generated"

	"strconv"

	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func GetTaskByIdController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	taskId := &pb.TaskId{Id: id}
	client, err := grpcclient.TaskManagementClient()
	if err != nil {
		http.Error(w, "Failed to connect to gRPC service", http.StatusInternalServerError)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.GetTaskById(ctx, taskId)
	if err != nil {
		log.Fatalf("Server error: %v", err)
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
