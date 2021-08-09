package main

import  (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
	"math/rand"
	"time"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "alex"
    password = "alexpass"
    dbname   = "alex"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"

var database *sql.DB

func CheckError(err error) {
    if err != nil {
        panic(err)
    }
}

func CreateRandomUrl() []byte {
        rand.Seed(time.Now().UnixNano())
	shorturl := make([]byte, 10)
	for i := range shorturl {
		shorturl[i] = letters[rand.Intn(len(letters))]
	}
	fmt.Println(shorturl)
	return shorturl
}

func InsertUrlToDB(longurl, shorturl string ) {
	insertStmt := `insert into "urls"(longurl, shorturl) values('`+longurl+`','`+shorturl+`')`
	fmt.Println(insertStmt)
        _, e := database.Exec(insertStmt)
        CheckError(e) 
}

func SelectLongUrlFromDB(longurl string) {
	rows, err := database.Query(`SELECT "shorturl" FROM "urls" WHERE longurl ='`+longurl+`'`)
	CheckError(err)
 
	defer rows.Close()
	for rows.Next() {
	var name string
	//var roll string
 
	err = rows.Scan(&name)
	CheckError(err)
 
	fmt.Println(name)
	}
 
	CheckError(err)
}

func main () {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
         
        // open database
    	db, err := sql.Open("postgres", psqlconn)
    	CheckError(err)
        
	database = db

	// close database
    	defer db.Close()
 
        // check db
    	err = db.Ping()
    	CheckError(err)
	var longurl string = "afsdsfsfafsad"
	var shorturl string = string(CreateRandomUrl())
	InsertUrlToDB(longurl,shorturl)
	SelectLongUrlFromDB(longurl)
//	insertStmt := `insert into "urs"("longurl", "shorturl") values('John', 'KJKJH')`
//	_, e := db.Exec(insertStmt)
//	CheckError(e) 
    	fmt.Println("Connected!")
}
