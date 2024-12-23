package routers

import (
	"github/http-server/controllers"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/task/{id}", controllers.GetTaskByIdController).Methods("GET")
	router.HandleFunc("/task", controllers.CreateTaskController).Methods("POST")
	return router
}
