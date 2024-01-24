//server/ main.go
package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	blog "github.com/bablu94/blog-service/proto"
	"google.golang.org/grpc"
)

type BlogServer struct {
	mu    sync.Mutex
	Posts map[string]*blog.PostResponse
	blog.UnimplementedBlogServiceServer
}

func (s *BlogServer) CreatePost(ctx context.Context, req *blog.PostRequest) (*blog.PostResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	postID := fmt.Sprintf("%d", len(s.Posts)+1)
	post := &blog.PostResponse{
		PostId:          postID,
		Title:           req.Title,
		Content:         req.Content,
		Author:          req.Author,
		PublicationDate: req.PublicationDate,
		Tags:            req.Tags,
	}

	s.Posts[postID] = post
	return post, nil
}

func (s *BlogServer) ReadPost(ctx context.Context, req *blog.PostIdRequest) (*blog.PostResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	post, ok := s.Posts[req.PostId]
	if !ok {
		return nil, fmt.Errorf("post not found")
	}

	return post, nil
}

func (s *BlogServer) UpdatePost(ctx context.Context, req *blog.UpdatePostRequest) (*blog.PostResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	post, ok := s.Posts[req.PostId]
	if !ok {
		return nil, fmt.Errorf("post not found")
	}

	post.Title = req.Title
	post.Content = req.Content
	post.Author = req.Author
	post.Tags = req.Tags

	return post, nil
}

func (s *BlogServer) DeletePost(ctx context.Context, req *blog.PostIdRequest) (*blog.DeleteResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.Posts[req.PostId]
	if !ok {
		return &blog.DeleteResponse{Success: false, ErrorMessage: "post not found"}, nil
	}

	delete(s.Posts, req.PostId)
	return &blog.DeleteResponse{Success: true}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	blogServer := &BlogServer{Posts: make(map[string]*blog.PostResponse)}
	blog.RegisterBlogServiceServer(server, blogServer)

	log.Println("Server is running on port 50051...")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
