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
	userRouter.HandleFunc("/refresh-token", controllers.RefreshTokenController).Methods("GET")
	authenticatedRouter := userRouter.PathPrefix("").Subrouter()
	authenticatedRouter.Use(middleware.AuthMiddleware)

	authenticatedRouter.HandleFunc("/task/{id}", controllers.GetTaskByIdController).Methods("GET")
	authenticatedRouter.HandleFunc("/tasks", controllers.GetTasksController).Methods("GET")
	authenticatedRouter.HandleFunc("/tasks", controllers.CreateTaskListController).Methods("POST")
	authenticatedRouter.HandleFunc("/task", controllers.CreateTaskController).Methods("POST")
	authenticatedRouter.HandleFunc("/tasks", controllers.CreateTasksController).Methods("POST")
	authenticatedRouter.HandleFunc("/task/{id}", controllers.DeleteTaskByIdController).Methods("DELETE")
	authenticatedRouter.HandleFunc("/task/{id}", controllers.UpdateTaskByIdController).Methods("PUT")

	authenticatedRouter.HandleFunc("/task/sort/{sortBy}", controllers.SortTasksControllers).Methods("GET")
	// authenticatedRouter.HandleFunc("/task/filter/{filterBy}", controllers.FilterTaskController).Methods("GET")
	return router
}
