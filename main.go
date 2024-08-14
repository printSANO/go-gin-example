package main

import (
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/printSANO/go-gin-example/database"
	"github.com/printSANO/go-gin-example/logger"
)

func main() {
	fmt.Println("Hello, World!")
	logger.Clog.Println("Hello, Debug!")
	database.Init("pgx")
}
