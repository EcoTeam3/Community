package main

import (
	"community/config"
	pb "community/generated/community"
	"community/service"
	"community/storage"
	"community/storage/postgres"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main(){
	db, err := storage.Connect()
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()

	cfg := config.Load()
	listener, err := net.Listen("tcp", ":" + cfg.URL_PORT)
	if err != nil{
		log.Fatal(err)
	}
	defer listener.Close()
	s := service.NewServerRepo(*postgres.NewCommunityRepo(db))
	grpc := grpc.NewServer()
	pb.RegisterCommunityServiceServer(grpc, s)
	if err = grpc.Serve(listener); err != nil{
		log.Fatal(err)
	}
}