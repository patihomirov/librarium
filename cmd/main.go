package main

import (
	"google.golang.org/grpc"
	"librarium/api"
	"librarium/server"
	"log"
	"net"
)

var ServiceOnline bool //Эта переменная нужна для функционального теста всего сервиса для определения что сервис запущен и не надо запускать его повторно

func main() {
	ServiceOnline = true
	log.Print("Service librarium is started at port 8080")

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
