package main

import (
	"coding-test-11-01-2025-dating-app/internal/handler"
	"coding-test-11-01-2025-dating-app/internal/middleware"
	"coding-test-11-01-2025-dating-app/internal/repository"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db := setupDatabase()
	repository := repository.NewRepository(db)
	handler := handler.NewServer(handler.NewServerOptions{
		Repository: repository,
	})
	router := setupRouter(handler)

	router.Run(":8081")
}

func setupDatabase() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	return db
}

func setupRouter(handler *handler.Server) *gin.Engine {
	router := gin.Default()

	// Public routes
	router.POST("/api/register", handler.Register)
	router.POST("/api/login", handler.Login)

	// Protected routes
	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	// it should be create profile page after register, but i will not make it
	// protected.POST("/profile", profileHandler.CreateProfile)
	protected.GET("/profiles", handler.GetProfilesToSwipe)
	protected.POST("/swipe", handler.Swipe)

	// ADMIN should be protected using another auth, but for this i'll not make it
	router.POST("/admin/verify-profile", handler.VerifyProfile)

	return router
}
