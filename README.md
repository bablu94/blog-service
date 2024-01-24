# blog gRPC Server and Client

implemented a blog gRPC server and client with the following functionality:
- create post
- read post
- update post
- delete post


# Running the application

1. Install the dependencies

```bash
go mod tidy
```

2. Run the server

```bash
go run server/main.go
```

3. Run the client

```bash
go run client/main.go