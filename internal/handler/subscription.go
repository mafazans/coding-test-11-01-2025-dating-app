package handler

import (
	"coding-test-11-01-2025-dating-app/internal/model"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Purchase premium subscription
func (h *Server) PurchasePremium(c *gin.Context) {
	userID := c.GetUint("userID")

	// In a real application, you would handle payment processing here

	// Create subscription for 30 days
	subscription := &model.Subscription{
		UserID:    userID,
		Status:    model.SubscriptionActive,
		StartDate: time.Now(),
		EndDate:   time.Now().AddDate(0, 1, 0), // 1 month subscription
	}

	if err := h.Repository.CreateSubscription(context.Background(), subscription); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create subscription"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Premium subscription activated successfully",
		"subscription": subscription,
	})
}

// Get subscription status
func (h *Server) GetSubscriptionStatus(c *gin.Context) {
	userID := c.GetUint("userID")

	subscription, err := h.Repository.GetActiveSubscription(context.Background(), userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No active subscription found"})
		return
	}

	c.JSON(http.StatusOK, subscription)
}

// Verify profile (could be an admin endpoint)
// should be receive callback after payment
func (h *Server) VerifyProfile(c *gin.Context) {
	var req struct {
		UserID uint `json:"user_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// for simplycity, ill just add the create subscription here
	subscription := &model.Subscription{
		UserID:    req.UserID,
		Status:    model.SubscriptionActive,
		StartDate: time.Now(),
		EndDate:   time.Now().AddDate(0, 1, 0),
	}

	if err := h.Repository.CreateSubscription(context.Background(), subscription); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create subscription"})
		return
	}

	if err := h.Repository.VerifyProfile(context.Background(), req.UserID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile verified successfully"})
}
