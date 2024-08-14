package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/printSANO/go-gin-example/config"
)

// 데이터베이스 연결을 수행하는 함수입니다.
func connectSQLDB(dbDriver string) (*sql.DB, error) {
	urlDB := config.DBURL
	db, err := sql.Open(dbDriver, urlDB)
	if err != nil {
		log.Fatalf("데이터베이스에 연결을 실패했습니다. 오류: %v\n", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// NewSQLDB 함수는 SQL 데이터베이스를 연결하고 DB 객체를 반환합니다.
func NewSQLDB(dbDriver string) (*DB, error) {
	db, err := connectSQLDB(dbDriver)
	if err != nil {
		return nil, err
	}
	log.Println("데이터베이스에 성공적으로 연결되었습니다.")
	return &DB{db}, nil
}

// ExecuteSQLFile 함수는 SQL 파일을 읽어서 SQL 명령을 실행합니다.
func ExecuteSQLFile(filePath string) error {
	sqlContent, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("SQL 파일을 읽는데 실패했습니다: %v", err)
	}
	commands := strings.Split(string(sqlContent), ";")

	for _, command := range commands {
		command = strings.TrimSpace(command)
		if command == "" {
			continue
		}

		_, err := SQLDB.Exec(command)
		if err != nil {
			return fmt.Errorf("SQL 명령을 실행하는데 실패했습니다: %v", err)
		}
	}

	return nil
}