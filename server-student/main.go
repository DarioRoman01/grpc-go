package main

import (
	"log"
	"net"

	"github.com/DarioRoman01/grpc/database"
	"github.com/DarioRoman01/grpc/server"
	"github.com/DarioRoman01/grpc/studentpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close()
	repo, err := database.NewPostgresRepository("postgres://postgres:postgres@localhost:54321/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	server := server.NewStudentsServer(repo)
	s := grpc.NewServer()
	studentpb.RegisterStudentServiceServer(s, server)

	reflection.Register(s)
	if err := s.Serve(listener); err != nil {
		log.Fatal(err)
	}

	log.Println("Server started on port 8080")
}
