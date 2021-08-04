package main

import (
	"grpcshorter/pkg/shorter"
	"google.golang.org/grpc"
	"grpcshorter/pkg/api"
	"net"
	"log"
)

func main() {
	s := grpc.NewServer()
	srv := &shorter.GRPCServer{}
	api.RegisterShorterServer(s, srv)
	
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
