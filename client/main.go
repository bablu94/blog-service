// client/main.go
package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	blog "github.com/bablu94/blog-service/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	client := blog.NewBlogServiceClient(conn)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		printMenu()
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			createPost(client, scanner)
		case "2":
			readPost(client, scanner)
		case "3":
			updatePost(client, scanner)
		case "4":
			deletePost(client, scanner)
		case "5":
			fmt.Println("Exiting the program.")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func printMenu() {
	fmt.Println("Choose an option:")
	fmt.Println("1. Create Post")
	fmt.Println("2. Read Post")
	fmt.Println("3. Update Post")
	fmt.Println("4. Delete Post")
	fmt.Println("5. Exit")
	fmt.Print("Enter your choice: ")
}

func createPost(client blog.BlogServiceClient, scanner *bufio.Scanner) {
	fmt.Print("Enter title for a new post: ")
	title := getUserInput(scanner)

	fmt.Print("Enter content for the post: ")
	content := getUserInput(scanner)

	fmt.Print("Enter author for the post: ")
	author := getUserInput(scanner)

	fmt.Print("Enter tags (comma-separated): ")
	tagsInput := getUserInput(scanner)
	tags := strings.Split(tagsInput, ",")

	createPostResponse, err := client.CreatePost(context.TODO(), &blog.PostRequest{
		Title:           title,
		Content:         content,
		Author:          author,
		PublicationDate: time.Now().Format(time.RFC3339),
		Tags:            tags,
	})
	if err != nil {
		log.Printf("Error creating post: %v", err)
	} else {
		log.Printf("Create Post Response: %+v", createPostResponse)
	}
}

func readPost(client blog.BlogServiceClient, scanner *bufio.Scanner) {
	fmt.Print("Enter post ID to read: ")
	postID := getUserInput(scanner)

	readPostResponse, err := client.ReadPost(context.TODO(), &blog.PostIdRequest{
		PostId: postID,
	})
	if err != nil {
		log.Printf("Error reading post: %v", err)
	} else {
		log.Printf("Read Post Response: %+v", readPostResponse)
	}
}

func updatePost(client blog.BlogServiceClient, scanner *bufio.Scanner) {
	fmt.Print("Enter post ID to update: ")
	postID := getUserInput(scanner)

	fmt.Print("Enter new title: ")
	title := getUserInput(scanner)

	fmt.Print("Enter new content: ")
	content := getUserInput(scanner)

	fmt.Print("Enter new author: ")
	author := getUserInput(scanner)

	fmt.Print("Enter new tags (comma-separated): ")
	tagsInput := getUserInput(scanner)
	tags := strings.Split(tagsInput, ",")

	updatePostResponse, err := client.UpdatePost(context.TODO(), &blog.UpdatePostRequest{
		PostId:  postID,
		Title:   title,
		Content: content,
		Author:  author,
		Tags:    tags,
	})
	if err != nil {
		log.Printf("Error updating post: %v", err)
	} else {
		log.Printf("Update Post Response: %+v", updatePostResponse)
	}
}

func deletePost(client blog.BlogServiceClient, scanner *bufio.Scanner) {
	fmt.Print("Enter post ID to delete: ")
	postID := getUserInput(scanner)

	deletePostResponse, err := client.DeletePost(context.TODO(), &blog.PostIdRequest{
		PostId: postID,
	})
	if err != nil {
		log.Printf("Error deleting post: %v", err)
	} else {
		log.Printf("Delete Post Response: %+v", deletePostResponse)
	}
}

func getUserInput(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}
