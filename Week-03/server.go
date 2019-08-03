package main

import (
	"./passenger"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	passenger.RegisterPassengerServiceServer(s, &passenger.PassengerFeedbackServerImp{
		FeedbackMap: make(map[string]*passenger.PassengerFeedback),
	})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}