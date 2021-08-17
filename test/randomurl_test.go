package postgres

import (
	"testing"
	"regexp"
	ru "github.com/ant1freeze/grpcshortener/shorter_server/randomurl"
)

func TestRandomUrl(t *testing.T) {
	n := 10
	want := regexp.MustCompile(`[a-zA-Z_]{10}`)
	randomurl := ru.CreateRandomUrl(n)
	if !want.MatchString(randomurl) {
		t.Fatalf(`ru.CreateRandomUrlUrl(10) = %q, want match for %#q, nil`, randomurl, want)
	}

}
