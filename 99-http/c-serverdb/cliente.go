package main

//go get -u github.com/go-sql-driver/mysql - para baixar a biblioteca do github

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/subosito/gotenv"
)

// Usuario

type Usuario struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

func init() {
	err := gotenv.Load(filepath.Join("C:/Users/Abner Silva/Documents/GitHub/Estudos_GO_LANG/", ".env"))
	if err != nil {
		panic(err)
	}

}

// UsuarioHandler analisa o request e delega para função adequada
func UsuarioHandler(res http.ResponseWriter, req *http.Request) {
	stringId := strings.TrimPrefix(req.URL.Path, "/usuario/")
	id, _ := strconv.Atoi(stringId)

	switch {
	case req.Method == "GET" && id > 0:
		usuarioPorID(res, req, id)
	case req.Method == "GET":
		usuarioTodos(res, req)
	default:
		res.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(res, "Sorry...")

	}

}

func usuarioPorID(res http.ResponseWriter, req *http.Request, id int) {

	login := os.Getenv("ROOT_PASSWORD_TABLE")

	// login = root:password
	db, err := sql.Open("mysql", login)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var u Usuario
	db.QueryRow("select id, nome from usuarios where id = ?", id).Scan(&u.ID, &u.Nome)

	json, _ := json.Marshal(u)
	fmt.Fprintln(res, string(json))

}

func usuarioTodos(res http.ResponseWriter, req *http.Request) {

	login := os.Getenv("ROOT_PASSWORD_TABLE")

	// login = root:password
	db, err := sql.Open("mysql", login)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	rows, _ := db.Query("select * from usuarios")
	defer rows.Close()

	var usuarios []Usuario
	for rows.Next() {
		var usuario Usuario
		rows.Scan(&usuario.ID, &usuario.Nome)
		usuarios = append(usuarios, usuario)

	}

	json, _ := json.Marshal(usuarios)

	fmt.Fprint(res, string(json))

}
