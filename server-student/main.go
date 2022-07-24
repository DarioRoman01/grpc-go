package main

import (
	"log"
	"net"

	"github.com/DarioRoman01/grpc/database"
	"github.com/DarioRoman01/grpc/server"
	"github.com/DarioRoman01/grpc/studentpb"
)

func main() {
	list, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	repo, err := database.NewPostgresRepository("postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	server := server.NewServer(repo)
	studentpb.RegisterStudentServiceServer(server, server)
}
