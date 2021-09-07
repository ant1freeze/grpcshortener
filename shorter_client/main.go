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

// Package main implements a client for GRPc URL Shortener service.
package main

import (
	"context"
	"log"
	"os"
	"time"
	
	"google.golang.org/grpc"
	"github.com/ant1freeze/grpcshortener/configs"
	pb "github.com/ant1freeze/grpcshortener"
)

var	cfg config.Config


func main() {
	conf, err := config.LoadConfig(".")//"$HOME/go/src/github.com/ant1freeze/grpcshortener/configs")
        if err != nil {
                log.Fatal("Can't get config from env file", err)
        }
	var serverAddr = conf.DBHost+conf.HttpPort
	// Set up a connection to the server.
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(5 * time.Second))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUrlShortenerClient(conn)

	// Contact the server and print out its response.
	var url, method string
	if len(os.Args) >= 3 {
		method = os.Args[1]
		url = os.Args[2]
	} else {
		method = "testMethod"
		url = "testUrl"
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// Check what method is in request
	if method == "create" {
		r, err := c.CreateUrl(ctx, &pb.UrlRequest{Urlreq: url}) //receive long url, find in db or create short url, insert it in db and return it 
		if err != nil {
			log.Fatalf("Can't create short url: %v", err)
		}
		log.Printf("localhost/%s",r.GetUrlrep())
	} else if method == "get" {
		r, err := c.GetUrl(ctx, &pb.UrlRequest{Urlreq: url}) //receive short url, find it in db and return long url or null if it not exist
                if err != nil {
                        log.Fatalf("Can't get long url: %v", err)
                }
		log.Printf("%s",r.GetUrlrep())
	} else if method == "testMethod" { //if Args < 3, print help
		log.Println("Need type 'get <short URL>' or 'create <long URL>'")
	} else {
		log.Println("Please choose one of two methods:\n-Get\n-Create\n\nafter method type Long or Short URL")
	}
}
