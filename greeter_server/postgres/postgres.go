package postgres

import  (
	"database/sql"
	_ "github.com/lib/pq"
)

var db *sql.DB

func Postgres (psqlconn string) (*sql.DB, error) {
//	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

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
