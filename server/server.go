package server

import (
	"fmt"

	"github.com/printSANO/go-gin-example/database"
	"github.com/printSANO/go-gin-example/handlers"
	"github.com/printSANO/go-gin-example/repositories"
	"github.com/printSANO/go-gin-example/services"
)

func Start(port string) {
	db := database.SetupDatabase()
	repo := repositories.NewRepository(db)
	service := services.NewService(repo)
	handler := handlers.NewHandler(service)

	router := setupRouter(handler)
	router.Run(fmt.Sprintf(":%s", port))
}
