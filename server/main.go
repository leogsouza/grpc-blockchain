package main

import (
	"log"
	"net"

	"github.com/leogsouza/grpc-blockchain/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalf("Unable to listen on 8080 port: %v", err)
	}

	srv := grpc.NewServer()
	proto.RegisterBlockchainServer(srv, &Server{})
	srv.Serve(listener)

}

type Server struct{}

func (s *Server) AddBlock(context.Context, *proto.AddBlockRequest) (*proto.AddBlockResponse, error) {
	return new(proto.AddBlockResponse), nil
}

func (s *Server) GetBlockchain(context.Context, *proto.GetBlockchainRequest) (*proto.GetBlockchainResponse, error) {
	return new(proto.GetBlockchainResponse), nil
}
