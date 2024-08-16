package main

import (
	"github.com/printSANO/go-gin-example/config"
	"github.com/printSANO/go-gin-example/server"
)

// 앱에서 가장 먼저 실행되는 함수
func main() {
	port := config.GetEnvVarAsString("PORT", "8080")
	server.Start(port)
}
