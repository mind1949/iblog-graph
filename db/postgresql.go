package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
)

var (
	pgDNS = "user=postgres user1dbname=iblog sslmode=disable"
	PGDB *sql.DB
)

func init() {
	PGDB, err := sql.Open("postgres", pgDNS)
	if err != nil {
		panic(fmt.Sprintf("sql.Open('postgresql', %q), err: %v", pgDNS, err))
	}

	err = PGDB.Ping()
	if err != nil {
		panic(fmt.Sprintf("err: PGDB.Ping(), %s", err))
	}
}
