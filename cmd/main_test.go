package main

import (
	"context"
	"fmt"
	"librarium/api"
	"librarium/client"
	"librarium/database_access"
	"librarium/server"
	"testing"
	"time"
)

func TestReadAuthorBook(t *testing.T) {
	//Arange
	bookauthor := "Грэм Грин"
	expected := []string{"Сила и слава"}
	//Act
	result := database_access.ReadAuthorBook(bookauthor)

	//Assert
	if fmt.Sprint(result) != fmt.Sprint(expected) {
		t.Errorf("Incorrest result. Expect %s, got %s", expected, result)
	}
}

func TestReadBookAuthor(t *testing.T) {
	//Arange
	bookname := "Волны гасят ветер"
	expected := []string{"Стругацкие"}
	//Act
	result := database_access.ReadBookAuthor(bookname)

	//Assert
	if fmt.Sprint(result) != fmt.Sprint(expected) {
		t.Errorf("Incorrest result. Expect %s, got %s", expected, result)
	}
}

func TestServerGetBooks(t *testing.T) {
	s := server.Server{}
	var req api.DataRequest
	//Arange
	req.AskMessage = "Грэм Грин"
	expected := `replyMessage:"Сила и слава"`

	//Act
	result, _ := s.GetBooks(context.Background(), &req)

	//Assert
	if fmt.Sprint(result) != fmt.Sprint(expected) {
		t.Errorf("Incorrest result. Expect %s, got %s", expected, result)
	}
}

func TestServerGetAuthor(t *testing.T) {
	s := server.Server{}
	var req api.DataRequest
	//Arange
	req.AskMessage = "Хроники Нарнии"
	expected := `replyMessage:"Льюис"`

	//Act
	result, _ := s.GetAuthor(context.Background(), &req)

	//Assert
	if fmt.Sprint(result) != fmt.Sprint(expected) {
		t.Errorf("Incorrest result. Expect %s, got %s", expected, result)
	}
}

func TestGRPCGetAuthor(t *testing.T) {
	//Arange
	bookName := "Волны гасят ветер"
	expected := []string{"Стругацкие"}
	//Act
	go main()
	time.Sleep(5 * time.Second)
	result := client.ClientGetAutor(bookName)

	//Assert
	if fmt.Sprint(result) != fmt.Sprint(expected) {
		t.Errorf("Incorrest result. Expect %s, got %s", expected, result)
	}

}

func TestGRPCGetBooks(t *testing.T) {
	//Arange
	bookAuthor := "Грэм Грин"
	expected := []string{"Сила и слава"}
	//Act
	go main()
	time.Sleep(5 * time.Second)
	result := client.ClientGetBooks(bookAuthor)

	//Assert
	if fmt.Sprint(result) != fmt.Sprint(expected) {
		t.Errorf("Incorrest result. Expect %s, got %s", expected, result)
	}

}
