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
	userRepo, profileRepo, swipeRepo, subscriptionRepo := initializeRepositories(db)
	authHandler, profileHandler, subscriptionHandler := initializeHandlers(userRepo, profileRepo, swipeRepo, subscriptionRepo)
	router := setupRouter(authHandler, profileHandler, subscriptionHandler)

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

func initializeRepositories(db *gorm.DB) (*repository.UserRepository, *repository.ProfileRepository, *repository.SwipeRepository, *repository.SubscriptionRepository) {
	userRepo := repository.NewUserRepository(db)
	profileRepo := repository.NewProfileRepository(db)
	swipeRepo := repository.NewSwipeRepository(db)
	subscriptionRepo := repository.NewSubscriptionRepository(db)
	return userRepo, profileRepo, swipeRepo, subscriptionRepo
}

func initializeHandlers(userRepo *repository.UserRepository, profileRepo *repository.ProfileRepository, swipeRepo *repository.SwipeRepository, subscriptionRepo *repository.SubscriptionRepository) (*handler.AuthHandler, *handler.ProfileHandler, *handler.SubscriptionHandler) {
	authHandler := handler.NewAuthHandler(userRepo)
	profileHandler := handler.NewProfileHandler(profileRepo, swipeRepo, subscriptionRepo)
	subscriptionHandler := handler.NewSubscriptionHandler(subscriptionRepo, profileRepo)
	return authHandler, profileHandler, subscriptionHandler
}

func setupRouter(authHandler *handler.AuthHandler, profileHandler *handler.ProfileHandler, subscriptionHandler *handler.SubscriptionHandler) *gin.Engine {
	router := gin.Default()

	// Public routes
	router.POST("/api/register", authHandler.Register)
	router.POST("/api/login", authHandler.Login)

	// Protected routes
	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	protected.GET("/profiles", profileHandler.GetProfilesToSwipe)

	// ADMIN should be protected using another auth, but for this i'll not make it
	router.POST("/admin/generate-profiles", profileHandler.GenerateDummyProfiles)
	router.POST("/admin/verify-profile", subscriptionHandler.VerifyProfile)

	return router
}
