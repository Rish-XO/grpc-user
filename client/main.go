package main

import (
    "context"
    "log"
    "time"

    "google.golang.org/grpc"
    pb "grpc-user-service/proto"
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := pb.NewUserServiceClient(conn)

    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    // Test GetUserById
    r, err := c.GetUserById(ctx, &pb.GetUserRequest{Id: 1})
    if err != nil {
        log.Fatalf("could not get user: %v", err)
    }
    log.Printf("User: %v", r.User)

    // Test GetUsersByIds
    r2, err := c.GetUsersByIds(ctx, &pb.GetUsersRequest{Ids: []int32{1, 2}})
    if err != nil {
        log.Fatalf("could not get users: %v", err)
    }
    log.Printf("Users: %v", r2.Users)

    // Test SearchUsers
    r3, err := c.SearchUsers(ctx, &pb.SearchUsersRequest{City: "LA"})
    if err != nil {
        log.Fatalf("could not search users: %v", err)
    }
    log.Printf("Users: %v", r3.Users)
}
