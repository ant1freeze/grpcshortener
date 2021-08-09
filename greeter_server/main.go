/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"log"
	"net"
	//"math/rand"
	//"time"
	"google.golang.org/grpc"
	pb "github.com/ant1freeze/grpcshortener"
	pg "github.com/ant1freeze/grpcshortener/greeter_server/postgres"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedUrlShortenerServer
}

//const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZs"


// SayHello implements helloworld.GreeterServer
func (s *server) Shorter(ctx context.Context, in *pb.UrlRequest) (*pb.UrlReply, error) {
	log.Printf("Received: %v", in.GetLongurl())
	log.Println(pg.RandomUrl())
	return &pb.UrlReply{Shorturl: "hello" + in.GetLongurl()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUrlShortenerServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
