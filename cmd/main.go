package main

import (
	"google.golang.org/grpc"
	"librarium/api"
	"librarium/server"
	"log"
	"net"
)

func main() {
	s := grpc.NewServer()
	srv := &server.Server{}
	api.RegisterLibrariumServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}

}
