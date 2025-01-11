package handler

import (
	"coding-test-11-01-2025-dating-app/internal/model"
	"coding-test-11-01-2025-dating-app/internal/repository"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type SubscriptionHandler struct {
	subscriptionRepo *repository.SubscriptionRepository
	profileRepo      *repository.ProfileRepository
}

func NewSubscriptionHandler(sr *repository.SubscriptionRepository, pr *repository.ProfileRepository) *SubscriptionHandler {
	return &SubscriptionHandler{
		subscriptionRepo: sr,
		profileRepo:      pr,
	}
}

// Purchase premium subscription
func (h *SubscriptionHandler) PurchasePremium(c *gin.Context) {
	userID := c.GetUint("userID")

	// In a real application, you would handle payment processing here

	// Create subscription for 30 days
	subscription := &model.Subscription{
		UserID:    userID,
		Status:    model.SubscriptionActive,
		StartDate: time.Now(),
		EndDate:   time.Now().AddDate(0, 1, 0), // 1 month subscription
	}

	if err := h.subscriptionRepo.CreateSubscription(subscription); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create subscription"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Premium subscription activated successfully",
		"subscription": subscription,
	})
}

// Get subscription status
func (h *SubscriptionHandler) GetSubscriptionStatus(c *gin.Context) {
	userID := c.GetUint("userID")

	subscription, err := h.subscriptionRepo.GetActiveSubscription(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No active subscription found"})
		return
	}

	c.JSON(http.StatusOK, subscription)
}

// Verify profile (could be an admin endpoint)
func (h *SubscriptionHandler) VerifyProfile(c *gin.Context) {
	var req struct {
		UserID uint `json:"user_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.profileRepo.VerifyProfile(req.UserID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile verified successfully"})
}
