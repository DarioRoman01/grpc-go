package main

import (
	"log"
	"net"

	"github.com/DarioRoman01/grpc/database"
	"github.com/DarioRoman01/grpc/server"
	"github.com/DarioRoman01/grpc/testpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}

	repo, err := database.NewPostgresRepository("postgres://postgres:postgres@localhost:54321/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	server := server.NewTestServer(repo)
	s := grpc.NewServer()
	testpb.RegisterTestServiceServer(s, server)

	reflection.Register(s)
	if err := s.Serve(listener); err != nil {
		log.Fatal(err)
	}

	log.Println("Server started on port 8080")
}
