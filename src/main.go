package main

import (
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"simple-microservice/crud"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load dotenv")
	}

	s := grpc.NewServer()
	srv := &crud.GRPCServer{}

	crud.RegisterCrudServer(s, srv)

	l, err := net.Listen("tcp", ":" + os.Getenv("grpc_port"))

	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
