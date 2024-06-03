package main

import (
    "context"
    "log"
    "net"

    "google.golang.org/grpc"
    pb "grpc-user-service/proto"
)

type User struct {
    ID      int32
    Fname   string
    City    string
    Phone   int64
    Height  float32
    Married bool
}

var users = []User{
    {ID: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
    {ID: 2, Fname: "John", City: "NY", Phone: 9876543210, Height: 5.9, Married: false},
    // Add more users as needed
}

type server struct {
    pb.UnimplementedUserServiceServer
}

func (s *server) GetUserById(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
    for _, user := range users {
        if user.ID == req.Id {
            return &pb.GetUserResponse{
                User: &pb.User{
                    Id:      user.ID,
                    Fname:   user.Fname,
                    City:    user.City,
                    Phone:   user.Phone,
                    Height:  user.Height,
                    Married: user.Married,
                },
            }, nil
        }
    }
    return nil, grpc.Errorf(codes.NotFound, "user not found")
}

func (s *server) GetUsersByIds(ctx context.Context, req *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
    var foundUsers []*pb.User
    for _, id := range req.Ids {
        for _, user := range users {
            if user.ID == id {
                foundUsers = append(foundUsers, &pb.User{
                    Id:      user.ID,
                    Fname:   user.Fname,
                    City:    user.City,
                    Phone:   user.Phone,
                    Height:  user.Height,
                    Married: user.Married,
                })
                break
            }
        }
    }
    return &pb.GetUsersResponse{Users: foundUsers}, nil
}

func (s *server) SearchUsers(ctx context.Context, req *pb.SearchUsersRequest) (*pb.SearchUsersResponse, error) {
    var foundUsers []*pb.User
    for _, user := range users {
        if (req.City == "" || user.City == req.City) &&
            (req.Phone == "" || user.Phone == req.Phone) &&
            (!req.Married || user.Married == req.Married) {
            foundUsers = append(foundUsers, &pb.User{
                Id:      user.ID,
                Fname:   user.Fname,
                City:    user.City,
                Phone:   user.Phone,
                Height:  user.Height,
                Married: user.Married,
            })
        }
    }
    return &pb.SearchUsersResponse{Users: foundUsers}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterUserServiceServer(s, &server{})
    log.Printf("server listening at %v", lis.Addr())
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
