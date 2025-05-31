package main

import (
	"log"
	"team-maker-api/config"
	"team-maker-api/delivery"
	"team-maker-api/delivery/handler"
	"team-maker-api/delivery/routes"
	"team-maker-api/repository"
	"team-maker-api/usecase"
	"team-maker-api/utils"
)

func main() {
	config.LoadConfig("config.json")
	utils.InitDB()

	userRepo := repository.NewUserRepository()
	authUsecase := usecase.NewAuthUsecase(userRepo)

	authHandler := handler.NewAuthHandler(authUsecase)
	// userHandler := handler.NewUserHandler()

	handlers := &delivery.Handlers{
		Auth: authHandler,
		// User: userHandler,
		// tambah handler lain di sini
	}

	r := routes.SetupRouter(handlers)

	log.Println("Server running at :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
