package main

import (
	"context"
	"log"
	"net"

	pb "github.com/lincolnjpg/grpc_go_course/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	pb.BlogServiceServer
}

var collection *mongo.Collection
var addr = "0.0.0.0:50051"

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("blogdb").Collection("blog")

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Could not listen: %v\n", err)
		return
	}

	log.Printf("Listenin on %v", addr)
	s := grpc.NewServer()
	pb.RegisterBlogServiceServer(s, &Server{})
	reflection.Register(s)

	if err = s.Serve(listener); err != nil {
		log.Fatalf("Could not serve: %v\n", err)
	}
}
