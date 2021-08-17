package postgres

import  (
	"math/rand"
	"time"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"

func CreateRandomUrl(i int) string {
        rand.Seed(time.Now().UnixNano())
	shorturl := make([]byte, i)
	for i := range shorturl {
		shorturl[i] = letters[rand.Intn(len(letters))]
	}
	return string(shorturl)
}

