package main

import (
	"github/http-server/routers"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const (
	PORT = ":8081"
)

func main() {
	router := routers.SetupRouter()
	log.Println("APP started and running on PORT " + PORT)
	err := http.ListenAndServe(PORT, router)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
