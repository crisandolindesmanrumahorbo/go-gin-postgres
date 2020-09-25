package db

import "fmt"

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "belajar_golang"
)

func Config() string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"sslmode=disable dbname=%s password=%s",
		host, port, user, dbname, password)
}
