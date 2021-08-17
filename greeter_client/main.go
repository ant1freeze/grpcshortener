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

// Package main implements a client for Greeter service.
package main

import (
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	pb "github.com/ant1freeze/grpcshortener"
)

const (
	address     = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
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
	if method == "create" {
		r, err := c.CreateUrl(ctx, &pb.UrlRequest{Urlreq: url})
		if err != nil {
			log.Fatalf("Can't create short url: %v", err)
		}
		log.Printf("localhost/%s",r.GetUrlrep())
	} else if method == "get" {
		r, err := c.GetUrl(ctx, &pb.UrlRequest{Urlreq: url})
                if err != nil {
                        log.Fatalf("Can't get long url: %v", err)
                }
		log.Printf("%s",r.GetUrlrep())
	} else if method == "testMethod" {
		log.Println("Need type 'get <short URL>' or 'create <long URL>'")
	} else {
		log.Println("Please choose one of two methods:\n-Get\n-Create\n\nafter method type Long or Short URL")
	}
}
