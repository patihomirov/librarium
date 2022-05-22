package database_access

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
)

func Read_book_autor(book_name string) string {
	db, err := sql.Open("mysql", "web:123@tcp(127.0.0.1:3306)/librarium_base")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	//	result, _ := db.Query("SELECT book_autor FROM librarium_base WHERE book_name = ?")

	//book_name = "Сила и слава"
	rows, _ := db.Query("SELECT book_autor FROM librarium_base WHERE book_name = ?", book_name)
	var book_autors []string
	var answer string

	for rows.Next() {
		rows.Scan(&answer)
		book_autors = append(book_autors, answer)
	}
	book_autor := fmt.Sprint(strings.Join(book_autors[:], "; "))
	log.Print("Автор запрошенной книги: " + book_autor)

	return book_autor
}
func Read_autor_book(book_autor string) string {
	db, err := sql.Open("mysql", "web:123@tcp(127.0.0.1:3306)/librarium_base")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	rows, _ := db.Query("SELECT book_name FROM librarium_base WHERE book_autor = ?", book_autor)
	var book_names []string
	var answer string
	for rows.Next() {
		rows.Scan(&answer)
		book_names = append(book_names, answer)
	}
	book_name := fmt.Sprint(strings.Join(book_names[:], "; "))
	log.Print("Книги запрошенного автора: " + book_name)
	return book_name
}
