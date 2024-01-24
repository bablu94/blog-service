// client/main_test.go
package main

import (
	"context"
	"testing"
	"time"

	blog "github.com/bablu94/blog-service/proto"
	"google.golang.org/grpc"
)

// MockBlogServiceServer is a mock implementation of the BlogServiceServer interface
type MockBlogServiceServer struct{}

func (s *MockBlogServiceServer) CreatePost(ctx context.Context, req *blog.PostRequest) (*blog.PostResponse, error) {

	return &blog.PostResponse{
		PostId:          "1",
		Title:           req.Title,
		Content:         req.Content,
		Author:          req.Author,
		PublicationDate: req.PublicationDate,
		Tags:            req.Tags,
	}, nil
}

func (s *MockBlogServiceServer) ReadPost(ctx context.Context, req *blog.PostIdRequest) (*blog.PostResponse, error) {

	return &blog.PostResponse{
		PostId:          req.PostId,
		Title:           "Mock Post",
		Content:         "This is a mock post.",
		Author:          "Mock Author",
		PublicationDate: time.Now().Format(time.RFC3339),
		Tags:            []string{"tag1", "tag2"},
	}, nil
}

func (s *MockBlogServiceServer) UpdatePost(ctx context.Context, req *blog.UpdatePostRequest) (*blog.PostResponse, error) {
	
	return &blog.PostResponse{
		PostId:          req.PostId,
		Title:           req.Title,
		Content:         req.Content,
		Author:          req.Author,
		PublicationDate: time.Now().Format(time.RFC3339),
		Tags:            req.Tags,
	}, nil
}

func (s *MockBlogServiceServer) DeletePost(ctx context.Context, req *blog.PostIdRequest) (*blog.DeleteResponse, error) {

	return &blog.DeleteResponse{
		Success: true,
	}, nil
}

// TestCreateReadUpdateDeletePost tests the CreatePost, ReadPost, UpdatePost, and DeletePost operations
func TestCreateReadUpdateDeletePost(t *testing.T) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()


	mockServer := &MockBlogServiceServer{}

	// Test CreatePost
	createPostResponse, err := mockServer.CreatePost(context.TODO(), &blog.PostRequest{
		Title:           "New Post",
		Content:         "This is the content of the new post.",
		Author:          "John Doe",
		PublicationDate: time.Now().Format(time.RFC3339),
		Tags:            []string{"tag1", "tag2"},
	})
	if err != nil {
		t.Fatalf("Error creating post: %v", err)
	}

	_ = createPostResponse

	// Test ReadPost
	readPostResponse, err := mockServer.ReadPost(context.TODO(), &blog.PostIdRequest{
		PostId: createPostResponse.PostId,
	})
	if err != nil {
		t.Fatalf("Error reading post: %v", err)
	}

	_ = readPostResponse

	// Test UpdatePost
	updatePostResponse, err := mockServer.UpdatePost(context.TODO(), &blog.UpdatePostRequest{
		PostId:  createPostResponse.PostId,
		Title:   "Updated Post",
		Content: "This is the updated content of the post.",
		Author:  "Jane Doe",
		Tags:    []string{"tag3", "tag4"},
	})
	if err != nil {
		t.Fatalf("Error updating post: %v", err)
	}

	_ = updatePostResponse

	// Test ReadPost (after update)
	readPostResponse, err = mockServer.ReadPost(context.TODO(), &blog.PostIdRequest{
		PostId: createPostResponse.PostId,
	})
	if err != nil {
		t.Fatalf("Error reading post after update: %v", err)
	}

	_ = readPostResponse

	// Test DeletePost
	deletePostResponse, err := mockServer.DeletePost(context.TODO(), &blog.PostIdRequest{
		PostId: createPostResponse.PostId,
	})
	if err != nil {
		t.Fatalf("Error deleting post: %v", err)
	}

	_ = deletePostResponse
}
