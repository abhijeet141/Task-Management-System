package controllers

import (
	"context"
	"encoding/json"
	grpcclient "github/http-server/grpc-client"
	pb "github/http-server/proto/generated"
	"log"
	"net/http"
	"time"
)

func CreateTasksController(w http.ResponseWriter, r *http.Request) {
	var tasks pb.TaskList
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
	stream, err := client.CreateTasks(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v", err)
	}
	for _, task := range tasks.Task {
		err := stream.Send(&pb.Task{
			UserId:      task.UserId,
			Title:       task.Title,
			Description: task.Description,
			Status:      task.Status,
		})
		if err != nil {
			log.Fatalf("Error while sending course details: %v", err)
		}
		log.Printf("Sent Task: %v", task)
		time.Sleep(1 * time.Second)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Printf("Error while receiving response: %v", err)
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
