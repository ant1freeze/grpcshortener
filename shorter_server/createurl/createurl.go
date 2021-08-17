/*
Package contains 2 func:
InsertUrl - for insert string with shorturl and longurl in DB
SelectShortUrl - for search shorturl in DB if we receive longurl; return shorturl
*/
package createurl

import  (
	"fmt"
	"database/sql"
)

func InsertUrl(longurl, shorturl string, database *sql.DB ) error {
	var insertStr string = fmt.Sprintf("insert into urls(longurl, shorturl) values('%s','%s')", longurl, shorturl)
	fmt.Printf(insertStr)
	_, err := database.Exec(insertStr)
        if err != nil {
		return err
	}
	return nil
}

func SelectShortUrl(longurl string, database *sql.DB) (shorturl string, err error) {
	var selectStr string = fmt.Sprintf("SELECT shorturl FROM urls WHERE longurl = '%s'", longurl)
	fmt.Printf(selectStr)
	rows, err := database.Query(selectStr)
	if err != nil {
		return "Can't make query 'SELECT shorturl' from urls table.", err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&shorturl)
		if err != nil {
			return "Can't make rows.Scan(&shorturl)", err
		}
	}
	return shorturl, err
}
