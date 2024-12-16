package usersDB

import (
	"database/sql"
	"os"
	"sync"
	"time"

	"github.com/Mateus-MS/CustomAuthInGolang/clusters"
)

type User struct {
	Id           string
	CreatedAt    time.Time
	Username     string
	Passwordhash string
	SessionToken string
	CsrfToken    string
}

var instance *sql.DB
var once sync.Once

func GetInstance() *sql.DB {
	once.Do(func() {
		instance = clusters.StartDBConnection(os.Getenv("DBuser"), os.Getenv("DBpass"), "users")
	})
	return instance
}
