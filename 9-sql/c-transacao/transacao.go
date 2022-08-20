package main

//go get -u github.com/go-sql-driver/mysql - para baixar a biblioteca do github

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	login := os.Getenv("ROOT_PASSWORD_TABLE")

	// login = root:password
	db, err := sql.Open("mysql", login)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into usuarios(id,nome) values(?,?)")

	stmt.Exec(2000, "bia")
	stmt.Exec(2001, "Carlos")

	_, err = stmt.Exec(1, "Tiago") // chave duplicada

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()

}
