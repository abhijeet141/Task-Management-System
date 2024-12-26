package routers

import (
	"github/http-server/controllers"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/task/{id}", controllers.GetTaskByIdController).Methods("GET")
	router.HandleFunc("/tasks", controllers.GetTasksController).Methods("GET")
	router.HandleFunc("/tasks", controllers.CreateTaskListController).Methods("POST")
	router.HandleFunc("/task", controllers.CreateTaskController).Methods("POST")
	router.HandleFunc("/tasks", controllers.CreateTasksController).Methods("POST")
	router.HandleFunc("/task/{id}", controllers.DeleteTaskByIdController).Methods("DELETE")
	router.HandleFunc("/task/{id}", controllers.UpdateTaskByIdController).Methods("PUT")
	return router
}
