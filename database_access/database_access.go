package database_access

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
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
	var book_autor string
	var book_autor_answer string
	for rows.Next() {
		rows.Scan(&book_autor_answer)
		book_autor = book_autor + book_autor_answer
	}
	log.Print("Автор запрошенной книги" + book_autor)
	return book_autor
}
