package main

import (
	"context"
	"fmt"
	"librarium/api/proto"
	"librarium/client"
	"librarium/database_access"
	"librarium/server"
	"testing"
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
	var req proto.DataRequest
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
	var req proto.DataRequest
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

//Для обоих следующих тестов требуется запуск самого сервиса main.
//Чтобы исключить одновременный запуск двух сервисов на одном порту, запретим запуск сервиса более одного раза.

func runMainService() {
	if ServiceOnline == false { //Запускает сервис если он еще не запущен
		go main()
		for ServiceOnline == true { //Не возвращаем управление в тест связи с сервисом пока сервис не запущен
		}
	}
}

func TestGRPCGetAuthor(t *testing.T) {
	//Arange
	bookName := "Волны гасят ветер"
	expected := []string{"Стругацкие"}
	//Act

	//go main()
	runMainService()
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
	//go main()
	runMainService()
	result := client.ClientGetBooks(bookAuthor)

	//Assert
	if fmt.Sprint(result) != fmt.Sprint(expected) {
		t.Errorf("Incorrest result. Expect %s, got %s", expected, result)
	}

}
