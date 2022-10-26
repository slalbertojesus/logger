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
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/slalbertojesus/logger/helloworld"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedLoggerServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) Log(ctx context.Context, in *pb.LogRequest) (*pb.LogReply, error) {
	log.Printf("Log Received at: [%v] \nLevel: [%v] \nMessage: [%v] \nDuration: [%v] \nMethod: [%v] \nProtocol: [%v] \nEnabled: [%v]",
		in.GetDateSent(),
		in.GetLevel(),
		in.GetMessage(),
		in.GetDuration(),
		in.GetMethod(),
		in.GetProtocol(),
		in.GetEnabled())
	//Here he would maybe add to Queue
	//We just log for now
	return &pb.LogReply{LogReply: "Message Received: " + in.GetMessage(), Ok: true}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterLoggerServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
