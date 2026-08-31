package main

import (
	"log"
	"online-ticketing/internal/config"
	"online-ticketing/internal/database"
	"online-ticketing/internal/handler"
	"online-ticketing/internal/middleware"
	"online-ticketing/internal/repository"
	"online-ticketing/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	db := database.ConnectDB(cfg)
	defer db.Close()

	r := gin.Default()
	ticketRepository := repository.NewTicketRepository(db)
	ticketService := service.NewTicketService(ticketRepository)
	ticketHandler := handler.NewTicketHandler(ticketService)

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository, []byte(cfg.SecretKey))
	userHandler := handler.NewUserHandler(userService)

	transactionRepository := repository.NewTransactionRepository(db)
	transactionService := service.NewTransactionService(transactionRepository)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	{
		protectedRoute := r.Group("/")
		protectedRoute.Use(middleware.AuthMiddleware([]byte(cfg.SecretKey)))
		protectedRoute.POST("/tickets", ticketHandler.BuyTicket)
		protectedRoute.POST("/tickets/create", ticketHandler.CreateTicket)
		protectedRoute.GET("/users", userHandler.GetUsers)
		protectedRoute.GET("/transactions", transactionHandler.GetTransaction)
	}

	{
		r.GET("/tickets", ticketHandler.GetTickets)
		r.POST("/users", userHandler.Register)
		r.POST("/login", userHandler.Login)
	}

	log.Printf("Server is running on port %s...", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
