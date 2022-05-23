package client

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"librarium/api/proto"
	"log"
)

func ClientGetAutor(bookName string) []string {
	conn, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	c := proto.NewLibrariumClient(conn)
	res, _ := c.GetAuthor(context.Background(), &proto.DataRequest{AskMessage: bookName})
	if err != nil {
		log.Fatal(err)
	}
	return res.GetReplyMessage()
}

func ClientGetBooks(bookAuthor string) []string {
	conn, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	c := proto.NewLibrariumClient(conn)
	res, _ := c.GetBooks(context.Background(), &proto.DataRequest{AskMessage: bookAuthor})
	if err != nil {
		log.Fatal(err)
	}
	return res.GetReplyMessage()
}
