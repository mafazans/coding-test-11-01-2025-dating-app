package handler

import (
	"coding-test-11-01-2025-dating-app/internal/model"
	"coding-test-11-01-2025-dating-app/internal/repository"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	profileRepo      *repository.ProfileRepository
	swipeRepo        *repository.SwipeRepository
	subscriptionRepo *repository.SubscriptionRepository
}

func NewProfileHandler(pr *repository.ProfileRepository, sr *repository.SwipeRepository, sbr *repository.SubscriptionRepository) *ProfileHandler {
	return &ProfileHandler{
		profileRepo:      pr,
		swipeRepo:        sr,
		subscriptionRepo: sbr,
	}
}

type SwipeRequest struct {
	ProfileID uint `json:"profile_id" binding:"required"`
	IsLike    bool `json:"is_like"`
}

func (h *ProfileHandler) GetProfilesToSwipe(c *gin.Context) {
	userID := c.GetUint("userID")

	// Check daily swipe limit
	count, err := h.swipeRepo.GetDailySwipeCount(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get swipe count"})
		return
	}

	if count >= 10 {
		c.JSON(http.StatusTooManyRequests, gin.H{"error": "Daily swipe limit reached"})
		return
	}

	profiles, err := h.profileRepo.GetUnswiped(userID, 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch profiles"})
		return
	}

	c.JSON(http.StatusOK, profiles)
}

func (h *ProfileHandler) Swipe(c *gin.Context) {
	userID := c.GetUint("userID")

	var req SwipeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if user is premium
	isPremium, err := h.subscriptionRepo.IsUserPremium(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check premium status"})
		return
	}

	// Only check swipe limit for non-premium users
	if !isPremium {
		count, err := h.swipeRepo.GetDailySwipeCount(userID)
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
		SwipedID: req.ProfileID,
		IsLike:   req.IsLike,
	}

	if err := h.swipeRepo.CreateSwipe(swipe); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to record swipe"})
		return
	}

	// Check for match if it's a like
	if req.IsLike {
		isMatch, err := h.swipeRepo.CheckMatch(userID, req.ProfileID)
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

// GenerateDummyProfiles creates dummy profiles for testing
func (h *ProfileHandler) GenerateDummyProfiles(c *gin.Context) {
	names := []string{"Alice", "Bob", "Charlie", "Diana", "Edward", "Fiona", "George", "Hannah", "Ian", "Julia"}
	genders := []string{"male", "female"}

	for i := 0; i < 500; i++ {
		profile := &model.Profile{
			Name:     names[rand.Intn(len(names))] + "-" + randomString(5),
			Bio:      "I love " + randomHobby(),
			Age:      18 + rand.Intn(42), // Ages 18-60
			Gender:   genders[rand.Intn(len(genders))],
			PhotoURL: fmt.Sprintf("https://placeholder.com/user%d.jpg", i),
		}

		if err := h.profileRepo.CreateProfile(profile); err != nil {
			continue
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Dummy profiles generated"})
}

func randomString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyz0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func randomHobby() string {
	hobbies := []string{"hiking", "reading", "photography", "cooking", "traveling", "music", "art", "sports", "gaming", "dancing"}
	return hobbies[rand.Intn(len(hobbies))]
}
