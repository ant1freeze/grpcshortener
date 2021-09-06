package main

import (
        "database/sql"
        _ "github.com/lib/pq"
        "fmt"
        "github.com/pressly/goose/v3"
        pg "github.com/ant1freeze/grpcshortener/internal/postgres"
        "github.com/ant1freeze/grpcshortener/configs"
)

var db *sql.DB
var conf config.Config

func main() {
        conf, err := config.LoadConfig(".")
        var psqlconn string = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",conf.DBUser, conf.DBPass, conf.DBHost, conf.DBPort, conf.DBName)
        db, err := sql.Open("postgres", psqlconn)
        if err != nil {
                panic(err)
        }
        db, err = pg.Postgres() //open and check db
        if err != nil {
                panic(err)
        }
        if err := goose.Up(db, "."); err != nil {
                panic(err)
        }
}
