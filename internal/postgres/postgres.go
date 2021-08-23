//Package contains 1 func
//Postgres - open DB and chect it is ok

package postgres

import  (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
	"log"
	"fmt"
	"github.com/ant1freeze/grpcshortener/configs"
)

var db *sql.DB
var cfg config.Config

func init() {
    // loads values from .env into the system
    if err := godotenv.Load(); err != nil {
        log.Print("No .env file found")
    }
}

func Postgres () (*sql.DB, error) {
        // open database
	conf := config.NewConfig()
	var psqlconn string = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",conf.DB.DBUser, conf.DB.DBPass, conf.DB.DBHost, conf.DB.DBPort, conf.DB.DBName)
	fmt.Println(psqlconn)
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
