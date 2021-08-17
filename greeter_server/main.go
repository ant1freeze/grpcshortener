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
	"fmt"
	//"math/rand"
	//"time"
	"database/sql"
	"google.golang.org/grpc"
	pb "github.com/ant1freeze/grpcshortener"
	pg "github.com/ant1freeze/grpcshortener/greeter_server/postgres"
	cr "github.com/ant1freeze/grpcshortener/greeter_server/createurl"
	get "github.com/ant1freeze/grpcshortener/greeter_server/geturl"
	ru "github.com/ant1freeze/grpcshortener/greeter_server/randomurl"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedUrlShortenerServer
}

var database *sql.DB

const (
    host     = "localhost"
    dbport     = 5432
    user     = "alex"
    password = "alexpass"
    dbname   = "alex"
)

var psqlconn string = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, dbport, user, password, dbname)

func (s *server) CreateUrl(ctx context.Context, in *pb.UrlRequest) (*pb.UrlReply, error) {
	log.Printf("\nMethod: Create\nReceived url: %v", in.GetUrlreq())
	database, err := pg.Postgres(psqlconn)
	if err != nil {
                return &pb.UrlReply{Urlrep: "error with db"}, err
        }

	shorturl, err := cr.SelectShortUrl(in.GetUrlreq(), database)
	if err != nil {
		return &pb.UrlReply{Urlrep: "error with SelectShortUrl"}, err
	}
	if shorturl != "" {
		return &pb.UrlReply{Urlrep: shorturl}, err
	} else {
	        shorturl = ru.CreateRandomUrl(10)
	        err := cr.InsertUrl(in.GetUrlreq(),shorturl,database)
		if err != nil {
			return &pb.UrlReply{Urlrep: "error with InsertUrl"}, err
		}
	}
//	if err != nil {
//		return &pb.UrlReply{Urlrep: shorturl}, err
//	}
	return &pb.UrlReply{Urlrep: shorturl}, nil
}

func (s *server) GetUrl(ctx context.Context, in *pb.UrlRequest) (*pb.UrlReply, error) {
	log.Printf("\nMethod: Get\nReceived url: %v", in.GetUrlreq())
	database, err := pg.Postgres(psqlconn)
	if err != nil {
		return &pb.UrlReply{Urlrep: "error with db"}, err
	}
	longurl, err := get.SelectLongUrl(in.GetUrlreq(), database)
        if err != nil {
                return &pb.UrlReply{Urlrep: longurl}, err
        }
        return &pb.UrlReply{Urlrep: longurl}, nil
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
