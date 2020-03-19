package mysql

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

// Info ...
type Info struct {
	Host string
	DB   string
	User string
	Pass string
	Loc  *time.Location
}

var (
	db  *sqlx.DB
	err error
)

// Connect ...
func (m Info) Connect() *sqlx.DB {
	// "user:password@/tpc(localhost:3306)/dbname?charset=utf8&parseTime=True&loc=Local"
	connStr := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=%s",
		m.User, m.Pass, m.Host, m.DB, m.Loc,
	)

	db, err = sqlx.Open("mysql", connStr)
	if err != nil {
		panic(err)
	}
	// defer db.Close()

	return db
}

// Close close the DB connection
func (m Info) Close() error {
	return db.Close()
}
