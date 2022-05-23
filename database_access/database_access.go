package database_access

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
)

func ReadBookAuthor(bookName string) []string {
	db, err := sql.Open("mysql", "web:123@tcp(127.0.0.1:3306)/librarium_base")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	//	result, _ := db.Query("SELECT book_autor FROM librarium_base WHERE book_name = ?")

	//book_name = "Сила и слава"
	rows, _ := db.Query("SELECT book_autor FROM librarium_base WHERE book_name = ?", bookName)
	var bookAuthors []string
	var answer string

	for rows.Next() {
		rows.Scan(&answer)
		bookAuthors = append(bookAuthors, answer)
	}
	bookAutor := fmt.Sprint(strings.Join(bookAuthors[:], "; "))
	log.Printf(`Автор запрошенной книги: "%s"`, bookAutor)

	return bookAuthors
}
func ReadAuthorBook(bookAutor string) []string {
	db, err := sql.Open("mysql", "web:123@tcp(127.0.0.1:3306)/librarium_base")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	rows, _ := db.Query("SELECT book_name FROM librarium_base WHERE book_autor = ?", bookAutor)
	var bookNames []string
	var answer string
	for rows.Next() {
		rows.Scan(&answer)
		bookNames = append(bookNames, answer)
	}
	bookName := fmt.Sprint(strings.Join(bookNames[:], "; "))
	log.Printf(`Книги запрошенного автора: "%s"`, bookName)
	return bookNames
}
