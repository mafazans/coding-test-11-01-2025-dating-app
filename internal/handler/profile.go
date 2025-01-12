package handler

import (
	"coding-test-11-01-2025-dating-app/internal/model"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SwipeRequest struct {
	UserID uint `json:"user_id" binding:"required"`
	IsLike bool `json:"is_like"`
}

func (h *Server) GetProfilesToSwipe(c *gin.Context) {
	userID := c.GetUint("userID")
	ctx := context.Background()

	_, err := h.Repository.GetUserByID(ctx, int(userID))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
		return
	}

	if limit <= 0 || limit > 100 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Limit must be between 1 and 100"})
		return
	}

	// Check if user is premium
	isPremium, err := h.Repository.IsUserPremium(ctx, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check premium status"})
		return
	}

	// Only check swipe limit for non-premium users
	if !isPremium {

		// Check daily swipe limit
		count, err := h.Repository.GetDailySwipeCount(ctx, userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get swipe count"})
			return
		}

		if count >= 10 {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Daily swipe limit reached"})
			return
		}

	}

	profiles, err := h.Repository.GetUnswiped(ctx, userID, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch profiles"})
		return
	}

	c.JSON(http.StatusOK, profiles)
}

func (h *Server) Swipe(c *gin.Context) {
	userID := c.GetUint("userID")

	ctx := context.Background()

	_, err := h.Repository.GetUserByID(ctx, int(userID))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	var req SwipeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if user is premium
	isPremium, err := h.Repository.IsUserPremium(ctx, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check premium status"})
		return
	}

	// Only check swipe limit for non-premium users
	if !isPremium {
		count, err := h.Repository.GetDailySwipeCount(ctx, userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get swipe count"})
			return
		}

		if count >= 10 {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error":   "Daily swipe limit reached",
				"message": "Upgrade to premium to remove the swipe limit!",
			})
			return
		}
	}

	// Record the swipe
	swipe := &model.Swipe{
		SwiperID: userID,
		SwipedID: req.UserID,
		IsLike:   req.IsLike,
	}

	_, err = h.Repository.GetUserByID(ctx, int(req.UserID))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	if err := h.Repository.CreateSwipe(ctx, swipe); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to record swipe"})
		return
	}

	// Check for match if it's a like
	if req.IsLike {
		isMatch, err := h.Repository.CheckMatch(ctx, userID, req.UserID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check match"})
			return
		}

		if isMatch {
			c.JSON(http.StatusOK, gin.H{
				"message":  "It's a match!",
				"is_match": true,
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Swipe recorded successfully",
		"is_match": false,
	})
}
