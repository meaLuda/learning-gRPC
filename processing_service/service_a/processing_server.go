package main

import (
	"context"
	"log"
	"net"

	pb "github.com/meaLuda/processingService/proto/processing"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedProcessingServiceServer
}

//🔥🔥 THIS SHOULD BE SAME NAME AS IN YOU PROTO RPC FUNCTION DEFINITION 🔥🔥🚀
func (s *server) ProcessRequest(ctx context.Context, in *pb.ProcessRequestMessage) (*pb.ProcessResponseMessage, error) {
	log.Println("⬆️ data incoming ")
	response := &pb.ProcessResponseMessage{
		RequestId: in.GetRequestId(),
		Result:    "Processed ☸️: " + in.GetData(),
	}
	return response, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProcessingServiceServer(s, &server{})
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
