package database

import (
	"database/sql"

	"github.com/printSANO/go-gin-example/logger"
)

var SQLDB *DB

type DB struct {
	*sql.DB
}

// 데이터베이스 연결 초기 함수
func Init(driverName string) {
	SQLDB, err := NewSQLDB(driverName)
	if err != nil {
		logger.Clog.Fatalln("데이터베이스에 연결을 실패했습니다.")
	}
	defer func() {
		if err = SQLDB.Close(); err != nil {
			panic(err)
		}
		logger.Clog.Println("데이터베이스에 연결이 끊겼습니다.")
	}()
}
