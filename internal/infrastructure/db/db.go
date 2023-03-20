package db

import (
	"fmt"
	_ "github.com/lib/pq"
)

func GetDsn(host, user, password, dbname, sslmode string, port uint16) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)
}
