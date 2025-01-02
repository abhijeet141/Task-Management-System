package controllers

// import (
// 	"context"
// 	grpcclient "github/http-server/grpc-client"
// 	pb "github/http-server/proto/generated"
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// func FilterTaskController(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	filterBy := vars["filterBy"]
// 	if filterBy != "Status" {
// 		http.Error(w, "Invalid filtering criteria", http.StatusBadRequest)
// 		return
// 	}
// 	filterTask := &pb.FilterTasksRequest{FilterBy: filterBy}
// 	client, err := grpcclient.TaskManagementClient()
// 	if err != nil {
// 		http.Error(w, "Failed to connect to gRPC service", http.StatusInternalServerError)
// 		return
// 	}
// 	res, err := client.FilterTasks(context.Background(), filterTask)
// 	if err != nil {
// 		log.Fatalf("Server error: %v", err)
// 	}
// }
