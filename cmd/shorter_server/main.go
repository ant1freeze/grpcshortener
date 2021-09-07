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

// Package main implements a server for GRPC URL Shortener service.
package main

import (
	"context"
	"database/sql"
	pb "github.com/ant1freeze/grpcshortener"
	cr "github.com/ant1freeze/grpcshortener/internal/createurl"
	get "github.com/ant1freeze/grpcshortener/internal/geturl"
	pg "github.com/ant1freeze/grpcshortener/internal/postgres"
	ru "github.com/ant1freeze/grpcshortener/internal/randomurl"
	"github.com/ant1freeze/grpcshortener/configs"
	"google.golang.org/grpc"
	"log"
	"net"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedUrlShortenerServer
}

var db *sql.DB
var cfg config.Config

func (s *server) CreateUrl(ctx context.Context, in *pb.UrlRequest) (*pb.UrlReply, error) {
	log.Printf("\nMethod: Create\nReceived url: %v", in.GetUrlreq())
	db, err := pg.Postgres()//open and check db
	if err != nil {
		return &pb.UrlReply{Urlrep: "error with db"}, err
	}

	shorturl, err := cr.SelectShortUrl(in.GetUrlreq(), db) //try find shorturl in DB
	if err != nil {
		return &pb.UrlReply{Urlrep: "error with SelectShortUrl"}, err
	}
	if shorturl != "" { //if we found shorturl, return it
		return &pb.UrlReply{Urlrep: shorturl}, err
	} else {
		shorturl = ru.CreateRandomUrl(10) //if didn't find, create random 10 letters
		err := cr.InsertUrl(in.GetUrlreq(), shorturl, db) //add new shorturl in db
		if err != nil {
			return &pb.UrlReply{Urlrep: "error with InsertUrl"}, err
		}
	}
	return &pb.UrlReply{Urlrep: shorturl}, nil
}

func (s *server) GetUrl(ctx context.Context, in *pb.UrlRequest) (*pb.UrlReply, error) {
	log.Printf("\nMethod: Get\nReceived url: %v", in.GetUrlreq())
	db, err := pg.Postgres() //open and check db
	if err != nil {
		return &pb.UrlReply{Urlrep: "error with db"}, err
	}
	longurl, err := get.SelectLongUrl(in.GetUrlreq(), db) //try find longurl in db
	if err != nil {
		return &pb.UrlReply{Urlrep: longurl}, err
	}
	return &pb.UrlReply{Urlrep: longurl}, nil
}

func main() {
	conf, err := config.LoadConfig(".")//"$HOME/go/src/github.com/ant1freeze/grpcshortener/configs")
        if err != nil {
                log.Fatal("Can't get config from env file", err)
        }

	lis, err := net.Listen("tcp", conf.HttpPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	log.Println("Server started...")
	pb.RegisterUrlShortenerServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
