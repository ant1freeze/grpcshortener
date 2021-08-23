//Package contains 1 func 
//SelectLongUrl - for search long url in DB if we receive shorturl; return longurl or text "Didn't find anything" if longurl doesn't exist

package geturl

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var db *sql.DB

func SelectLongUrl(shorturl string, database *sql.DB) (longurl string, err error) {
	selectStr := fmt.Sprintf("SELECT longurl FROM urls WHERE shorturl = '%s'", shorturl)
	rows, err := database.Query(selectStr)
	if err != nil {
		return "Can't make query 'SELECT longurl' from urls table.", err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&longurl)
		if err != nil {
			return "Can't make rows.Scan(longurl)", err
		}
	}
	if string(longurl) == "" {
		longurl = "Didn't find anything."
	}
	return string(longurl), err
}
