package db

import (
	"fmt"
)

func GetDsn(host, user, password, dbname, schema, sslmode string, port uint16) string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s search_path=%s",
		host, port, user, password, dbname, sslmode, schema,
	)
}
