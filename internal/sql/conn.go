package sql

import (
	"database/sql"

	_ "github.com/microsoft/go-mssqldb"
	"github.com/ronymmoura/adiq-recurrence-check/internal/util"
)

type DbConn struct {
	*sql.DB
}

func CreateConnection(config util.Config) (*DbConn, error) {
	db, err := sql.Open("sqlserver", config.DatabaseUrl)
	conn := &DbConn{
		db,
	}
	return conn, err
}
