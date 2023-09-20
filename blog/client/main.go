package main

import (
	"log"

	pb "github.com/lincolnjpg/grpc_go_course/blog/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr = "0.0.0.0:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect: %v\n", err)
		return
	}
	defer conn.Close()

	c := pb.NewBlogServiceClient(conn)
	id := createBlog(c)
	readBlog(c, id)
	readBlog(c, "aNonExistingID")
	updateBlog(c, id)
	listBlog(c)
	deleteBlog(c, id)
}
