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

	stmt, _ := db.Prepare("update usuarios set nome = ? where id = ?")
	stmt.Exec("Abner Silva", 1)

	stmt2, _ := db.Prepare("delete from usuarios where id = ?")
	stmt2.Exec(2)

}
