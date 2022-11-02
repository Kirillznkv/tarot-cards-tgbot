package sqlstore

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"strings"
	"testing"
)

func TestDB(t *testing.T, databaseURL string) (*sql.DB, func(...string)) {
	t.Helper()

	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		t.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}

	return db, func(tabels ...string) {
		defer db.Close()

		if len(tabels) > 0 {
			if _, err := db.Exec(fmt.Sprintf(("TRUNCATE %s CASCADE"), strings.Join(tabels, ", "))); err != nil {
				t.Fatal(err)
			}
		}
	}
}
