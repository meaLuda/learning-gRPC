package main

// API gateway that
// handles REST requests and communicates with the gRPC

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	pb "github.com/meaLuda/processingClientService/proto/processing"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Request struct {
	RequestID string `json:"request_id"`
	Data      string `json:"data"`
}

type Response struct {
	RequestID string `json:"request_id"`
	Result      string `json:"data"`
}



func main() {
	r := mux.NewRouter()
	r.HandleFunc("/process",processRequestHandler).Methods("POST")
	log.Println("API Gateway listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}


var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func processRequestHandler(w http.ResponseWriter, r *http.Request) {
	var req Request
	fmt.Println(r.Body)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Connect to grpc server and send data
	conn, err := grpc.NewClient(*addr,grpc.WithTransportCredentials(insecure.NewCredentials()) )	
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewProcessingServiceClient(conn)

	// Call gRPC service 
	grpcReq := &pb.ProcessRequestMessage{
		RequestId: req.RequestID,
    	Data:       req.Data,
	}

	// get response from our microservice 
	grpcResp, err := client.ProcessRequest(context.Background(), grpcReq)
	if err != nil {
		http.Error(w, err.Error(),http.StatusInternalServerError)
		return
	}
	fmt.Println(grpcResp)
    
	// Send response
	resp := &Response{
		RequestID: grpcResp.GetRequestId(),
		Result:    grpcResp.GetResult(),
	}
	fmt.Println(resp)

	w.Header().Set("Content-type","application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
