package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Connection struct {
	Dsn string
	Db  *sql.DB
}

func getDsn(host, user, password, dbname, sslmode string, port int) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)
}

func NewConnection(host, user, password, dbname, sslmode string, port int) (*Connection, error) {
	dsn := getDsn(host, user, password, dbname, sslmode, port)
	db, err := sql.Open("postgres", dsn)

	return &Connection{Db: db, Dsn: dsn}, err
}
