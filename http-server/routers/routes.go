package routers

import (
	"github/http-server/controllers"
	"github/http-server/middleware"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	userRouter := router.PathPrefix("/api/v1/user").Subrouter()

	userRouter.HandleFunc("/register", controllers.RegisterUserController).Methods("POST")
	userRouter.HandleFunc("/login", controllers.LoginUserController).Methods("POST")

	userRouter.Use(middleware.AuthMiddleware)

	userRouter.HandleFunc("/task/{id}", controllers.GetTaskByIdController).Methods("GET")
	userRouter.HandleFunc("/tasks", controllers.GetTasksController).Methods("GET")
	userRouter.HandleFunc("/tasks", controllers.CreateTaskListController).Methods("POST")
	userRouter.HandleFunc("/task", controllers.CreateTaskController).Methods("POST")
	userRouter.HandleFunc("/tasks", controllers.CreateTasksController).Methods("POST")
	userRouter.HandleFunc("/task/{id}", controllers.DeleteTaskByIdController).Methods("DELETE")
	userRouter.HandleFunc("/task/{id}", controllers.UpdateTaskByIdController).Methods("PUT")
	return router
}
