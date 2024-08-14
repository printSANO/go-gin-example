package main

import (
	"fmt"

	"github.com/printSANO/go-gin-example/logger"
)

func main() {
	fmt.Println("Hello, World!")
	logger.Clog.Println("Hello, Debug!")
}
