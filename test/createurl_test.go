package postgres

import (
	"testing"
	"regexp"
	"fmt"
	cr "github.com/ant1freeze/grpcshortener/greeter_server/createurl"
	"database/sql"
        _ "github.com/lib/pq"

)

var database *sql.DB

const (
        host     = "localhost"
        dbport   = 5432
        user     = "alex"
        password = "alexpass"
        dbname   = "alex"
)

var psqlconn string = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, dbport, user, password, dbname)

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

func TestSelectShortUrl(t *testing.T) {
	database, err := Postgres(psqlconn)
	defer database.Close()
//	longurl := "www.google.com"
	want := regexp.MustCompile(`[a-zA-Z_]{10}`)
	shorturl, err := cr.SelectShortUrl("www.google.com", database)
	fmt.Println(shorturl)
	if !want.MatchString(shorturl) || err != nil {
		t.Fatalf(`get.SelectShortUrl("www.google.com") = %q, %v want match for %#q, nil`, shorturl, err, want)
	}

}
