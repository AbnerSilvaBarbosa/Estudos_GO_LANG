package main

//go get -u github.com/go-sql-driver/mysql - para baixar a biblioteca do github

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type usuario struct {
	id   int
	nome string
}

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

	rows, _ := db.Query("select id,nome from usuarios where id > ?", 5)
	defer rows.Close()

	for rows.Next() {
		var u usuario
		rows.Scan(&u.id, &u.nome)
		fmt.Println(u)
	}
}
