package main

import (
	"fmt"
	"librarium/database_access"
	"testing"
)

func TestReadAuthorBook(t *testing.T) {
	//Arange
	bookautor := "Грэм Грин"
	expected := []string{"Сила и слава"}
	//Act
	result := database_access.ReadAuthorBook(bookautor)

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
