//Package contains 1 func
//Postgres - open DB and chect it is ok

package postgres

import  (
	"database/sql"
	_ "github.com/lib/pq"
)

var db *sql.DB

func Postgres (psqlconn string) (*sql.DB, error) {
        // open database
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
