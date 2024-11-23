// server/main.go
package main

import (
	"context"
	"log"
	"net"
	pb "grpc-tutorial/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"    // Add this import
	"google.golang.org/grpc/status"   // Add this import
)

type server struct {
	pb.UnimplementedUserServiceServer
	users map[int32]*pb.UserResponse
}

func newServer() *server {
	// Initialize with some dummy data
	s := &server{
		users: make(map[int32]*pb.UserResponse),
	}
	s.users[1] = &pb.UserResponse{UserId: 1, Name: "Alice", Email: "alice@example.com"}
	s.users[2] = &pb.UserResponse{UserId: 2, Name: "Bob", Email: "bob@example.com"}
	return s
}

func (s *server) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	if user, ok := s.users[req.UserId]; ok {
		return user, nil
	}
	return nil, status.Errorf(codes.NotFound, "user not found")
}

func (s *server) ListUsers(ctx context.Context, _ *pb.Empty) (*pb.UserList, error) {
	users := make([]*pb.UserResponse, 0, len(s.users))
	for _, user := range s.users {
		users = append(users, user)
	}
	return &pb.UserList{Users: users}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, newServer())
	
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}