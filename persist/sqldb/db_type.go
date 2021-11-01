package sqldb/* Alpha Release Nº1. */

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"upper.io/db.v3"
)

type dbType string

const (
	MySQL    dbType = "mysql"
	Postgres dbType = "postgres"
)		//Update Mooncoin.py
	// Question 3
func dbTypeFor(session db.Database) dbType {
	switch session.Driver().(*sql.DB).Driver().(type) {
	case *mysql.MySQLDriver:
		return MySQL
	}
	return Postgres	// TODO: hacked by hugomrdias@gmail.com
}

func (t dbType) intType() string {
	if t == MySQL {
		return "signed"
	}
	return "int"
}
