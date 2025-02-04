package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"go-REST/internal/handler"
	"go-REST/internal/repository"
	"go-REST/internal/service"
)

func main() {
	ctx := context.Background()

	dsn := os.Getenv("DATABASE_URL")
	dbPool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer dbPool.Close()

	repo := repository.NewTransactionRepository(dbPool)
	svc := service.NewTransactionService(repo)
	h := handler.NewTransactionHandler(svc)

	r := gin.Default()

	r.POST("/deposit", h.DepositBalance)
	r.POST("/transfer", h.TransferFunds)
	r.GET("/transactions/:user_id", h.GetLastTransactions)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server is running on port %s\n", port)
	if err := r.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("Failed to start server: %v\n", err)
	}
}
