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

func DeleteTaskByIdController(w http.ResponseWriter, r *http.Request) {
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
	res, err := client.DeleteTaskById(ctx, taskId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to delete task: %v", err), http.StatusInternalServerError)
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
