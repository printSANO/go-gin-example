package database

import (
	"database/sql"
	"fmt"

	"github.com/printSANO/go-gin-example/logger"
)

var SQLDB *DB

type DB struct {
	*sql.DB
}

func Init(driverName string) {
	SQLDB, err := NewSQLDB(driverName)
	if err != nil {
		logger.Clog.Println("데이터베이스에 연결을 실패했습니다.")
		panic(fmt.Errorf("데이터베이스 연결을 실패했습니다. 이유: \n%v", err))
	}
	defer func() {
		if err = SQLDB.Close(); err != nil {
			panic(err)
		}
		logger.Clog.Println("데이터베이스에 연결이 끊겼습니다.")
	}()
}
