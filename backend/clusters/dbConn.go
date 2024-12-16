package clusters

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func StartDBConnection(user, pass, dbname string) *sql.DB {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, pass, dbname)

	client, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to postgreSQL: ", err)
	}

	return client
}
