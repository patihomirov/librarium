package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"librarium/api/proto"
	"librarium/server"
	"log"
	"net"
)

var ServiceOnline bool //Эта переменная нужна для функционального теста всего сервиса для определения что сервис запущен и не надо запускать его повторно

func main() {
	ServiceOnline = true
	if selftests() {
		log.Print("Service librarium is started at port 8080")
	}

	s := grpc.NewServer()
	srv := &server.Server{}
	proto.RegisterLibrariumServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}

func selftests() bool {
	log.Println("selftesting is started")
	log.Println("mySQL access testing")

	s := server.Server{}
	var req proto.DataRequest
	//Arange
	req.AskMessage = "Грэм Грин"
	expected := `replyMessage:"Сила и слава"`

	//Act
	result, _ := s.GetBooks(context.Background(), &req)

	//Assert
	if fmt.Sprint(result) == fmt.Sprint(expected) {
		log.Println("testing is done")
		return true
	} else {
		log.Fatalf("testing is fail. Expect %s, got %s", expected, result)
		return false
	}

}
