package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/danush754/Task-Manager/internal/db"
	handler "github.com/danush754/Task-Manager/internal/handlers"
	"github.com/danush754/Task-Manager/internal/repository"
	"github.com/danush754/Task-Manager/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatalf("Error connecting DB: %v", "failed to  get db credentials")
	}

	newContext := context.Background()

	pool, err := db.ConnectDB(newContext, dsn)
	if err != nil {
		log.Fatalf("Error connecting to DB: %v", err)
	}

	defer pool.Close()

	queries := db.New(pool)

	userRepo := repository.NewUserRepo(queries)

	userService := service.NewUserService(userRepo)
	authHandler := handler.NewAuthHandler(userService)

	h := handler.NewHandler(queries)
	//  set up the server
	r := gin.Default()
	RegisterRoutes(r, h, authHandler)

	srvErr := make(chan error, 1)
	go func() {
		srvErr <- r.Run(":8080")
	}()

	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, os.Interrupt, syscall.SIGTERM)
	select {
	case sig := <-quitChannel:
		log.Printf("shutting down due to %v", sig)
		timeoutContext, cancel := context.WithTimeout(newContext, 5*time.Second)
		defer cancel()
		pool.Close()
		_ = timeoutContext
	case err := <-srvErr:
		log.Fatalf("server error %v:", err)
	}

}
