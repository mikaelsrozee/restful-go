package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type App struct {
    Router *mux.Router
    DB *sql.DB
}

func (a *App) Init(user string, password string, dbname string) {
    connString := fmt.Sprintf("%s:%s@/%s", user, password, dbname)
    db, err := sql.Open("mysql", connString)

    if err != nil {
        panic(err)
    }

    fmt.Println("App started", db)
}
