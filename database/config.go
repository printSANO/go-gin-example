package database

import "database/sql"

var SQLDB *DB

type DB struct {
	*sql.DB
}
