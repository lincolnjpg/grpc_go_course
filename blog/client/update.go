package main

import (
	"context"
	"log"

	pb "github.com/lincolnjpg/grpc_go_course/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("updateBlog was invoked")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "José",
		Title:    "Other title",
		Content:  "Some content",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)
	if err != nil {
		log.Fatalf("Error happened while updating: %v\n", err)
	}

	log.Println("Blog was updated")
}
