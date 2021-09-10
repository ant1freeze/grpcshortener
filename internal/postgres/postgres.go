//Package contains 1 func
//Postgres - open DB and chect it is ok

package postgres

import  (
	"database/sql"
	"log"
	"fmt"
	
	_ "github.com/lib/pq"
	"github.com/ant1freeze/grpcshortener/configs"
)

var db *sql.DB
var cfg config.Config

func Postgres () (*sql.DB, error) {
        // open database
	conf, err := config.LoadConfig(".")//"$HOME/go/src/github.com/ant1freeze/grpcshortener/configs")
	if err != nil {
		log.Fatal("Can't get config from env file", err)
	}
	var psqlconn string = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",conf.DBUser, conf.DBPass, conf.DBHost, conf.DBPort, conf.DBName)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return db, err
	}

        // check db
	err = db.Ping()
	if err != nil {
		return db, err
	}
	return db, err
}
