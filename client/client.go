package client

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"librarium/api"
	"log"
)

func ClientGetAutor(bookName string) []string {
	conn, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	c := api.NewLibrariumClient(conn)
	res, _ := c.GetAuthor(context.Background(), &api.DataRequest{AskMessage: bookName})
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

	c := api.NewLibrariumClient(conn)
	res, _ := c.GetBooks(context.Background(), &api.DataRequest{AskMessage: bookAuthor})
	if err != nil {
		log.Fatal(err)
	}
	return res.GetReplyMessage()
}
