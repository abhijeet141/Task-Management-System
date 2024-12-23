package main

import (
	pb "github/grpc-server/proto/generated"
	"os"
	"time"

	"log"
	"net"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

const (
	PORT = ":8080"
)

type TaskManagementServer struct {
	pb.TaskManagementServiceServer
}

type Task struct {
	Id          int       `orm:"column(id);auto"`
	Title       string    `orm:"column(title)"`
	Description string    `orm:"column(description)"`
	Status      string    `orm:"column(status)"`
	CreatedAt   time.Time `orm:"column(created_at);type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Task))
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	connectionstring := os.Getenv("CONNECTION_STRING")
	if connectionstring == "" {
		log.Fatal("Connection String is not set in the environment")
	}
	orm.RegisterDataBase("default", "mysql", connectionstring)

	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	log.Println("Database connection is established")
}

func main() {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("Failed to start the server %v", err)
	}
	log.Println("APP started and running on PORT 8080")
	grpcserver := grpc.NewServer()
	pb.RegisterTaskManagementServiceServer(grpcserver, &TaskManagementServer{})
	err = grpcserver.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
