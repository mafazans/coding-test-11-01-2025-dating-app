package handler_test

import (
	"coding-test-11-01-2025-dating-app/internal/handler"
	"coding-test-11-01-2025-dating-app/internal/model"
	"coding-test-11-01-2025-dating-app/internal/repository"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"golang.org/x/crypto/bcrypt"
)

func TestRegister(t *testing.T) {
	type testCase struct {
		name           string
		requestBody    string
		setupMock      func(*repository.MockRepositoryInterface)
		expectedStatus int
		expectedBody   string
	}

	tests := []testCase{
		{
			name:        "Valid registration request",
			requestBody: `{"username": "testuser@gmail.com", "password": "password123"}`,
			setupMock: func(mockRepo *repository.MockRepositoryInterface) {
				mockRepo.EXPECT().
					CreateUser(
						gomock.Any(),
						gomock.Any(),
					).Return(nil)
			},
			expectedStatus: http.StatusCreated,
			expectedBody:   "",
		},
		{
			name:           "Empty credentials",
			requestBody:    `{"username": "", "password": ""}`,
			setupMock:      func(mockRepo *repository.MockRepositoryInterface) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "error",
		},
		{
			name:           "Invalid username format",
			requestBody:    `{"username": "invalid username!", "password": "password123"}`,
			setupMock:      func(mockRepo *repository.MockRepositoryInterface) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "error",
		},
		{
			name:        "Failed to create user",
			requestBody: `{"username": "testuser@gmail.com", "password": "password123"}`,
			setupMock: func(mockRepo *repository.MockRepositoryInterface) {
				mockRepo.EXPECT().
					CreateUser(
						gomock.Any(),
						gomock.AssignableToTypeOf(&model.User{}),
					).
					DoAndReturn(func(_ context.Context, u *model.User) error {
						assert.Equal(t, "testuser@gmail.com", u.Username)
						return errors.New("database error")
					})
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "Failed to create user",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := repository.NewMockRepositoryInterface(ctrl)
			server := handler.NewServer(handler.NewServerOptions{Repository: mockRepo})

			// Configure mock expectations
			tc.setupMock(mockRepo)

			// Create test request
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = httptest.NewRequest("POST", "/register", strings.NewReader(tc.requestBody))
			c.Request.Header.Set("Content-Type", "application/json")

			// Execute test
			server.Register(c)

			// Assert results
			assert.Equal(t, tc.expectedStatus, w.Code)
			if tc.expectedBody != "" {
				assert.Contains(t, w.Body.String(), tc.expectedBody)
			}
		})
	}
}

func TestLogin(t *testing.T) {
	type testCase struct {
		name           string
		requestBody    string
		setupMock      func(*repository.MockRepositoryInterface)
		setupEnv       func()
		cleanupEnv     func()
		expectedStatus int
		expectedBody   string
	}

	// Helper function to create hashed password
	createHashedPassword := func(password string) string {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		return string(hashedPassword)
	}

	tests := []testCase{
		{
			name:        "Successful login",
			requestBody: `{"username": "testuser@gmail.com", "password": "password123"}`,
			setupMock: func(mockRepo *repository.MockRepositoryInterface) {
				mockRepo.EXPECT().
					GetUserByUsername(gomock.Any(), "testuser@gmail.com").
					Return(&model.User{
						Username: "testuser@gmail.com",
						Password: createHashedPassword("password123"),
					}, nil)
			},
			setupEnv: func() {
				os.Setenv("JWT_SECRET", "test-secret")
			},
			cleanupEnv: func() {
				os.Unsetenv("JWT_SECRET")
			},
			expectedStatus: http.StatusOK,
			expectedBody:   "token",
		},
		{
			name:           "Invalid JSON request",
			requestBody:    `{"username": "testuser", invalid json}`,
			setupMock:      func(mockRepo *repository.MockRepositoryInterface) {},
			setupEnv:       func() {},
			cleanupEnv:     func() {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "error",
		},
		{
			name:        "User not found",
			requestBody: `{"username": "testuser@gmail.com", "password": "password123"}`,
			setupMock: func(mockRepo *repository.MockRepositoryInterface) {
				mockRepo.EXPECT().
					GetUserByUsername(gomock.Any(), "testuser@gmail.com").
					Return(nil, errors.New("not found"))
			},
			setupEnv:       func() {},
			cleanupEnv:     func() {},
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   "Invalid credentials",
		},
		{
			name:        "Invalid password",
			requestBody: `{"username": "testuser@gmail.com", "password": "password123"}`,
			setupMock: func(mockRepo *repository.MockRepositoryInterface) {
				mockRepo.EXPECT().
					GetUserByUsername(gomock.Any(), "testuser@gmail.com").
					Return(&model.User{
						Username: "testuser@gmail.com",
						Password: createHashedPassword("different-password"),
					}, nil)
			},
			setupEnv: func() {
				os.Setenv("JWT_SECRET", "test-secret")
			},
			cleanupEnv: func() {
				os.Unsetenv("JWT_SECRET")
			},
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   "Invalid credentials",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := repository.NewMockRepositoryInterface(ctrl)
			server := handler.NewServer(handler.NewServerOptions{Repository: mockRepo})

			// Configure mock and environment
			tc.setupMock(mockRepo)
			tc.setupEnv()
			defer tc.cleanupEnv()

			// Create test request
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = httptest.NewRequest("POST", "/login", strings.NewReader(tc.requestBody))
			c.Request.Header.Set("Content-Type", "application/json")

			// Execute test
			server.Login(c)

			// Assert results
			assert.Equal(t, tc.expectedStatus, w.Code)
			if tc.expectedBody != "" {
				assert.Contains(t, w.Body.String(), tc.expectedBody)
			}
		})
	}
}

// Premium user successfully swipes right and creates a match
func TestPremiumUserSwipeRightCreatesMatch(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockRepositoryInterface(ctrl)
	server := handler.NewServer(handler.NewServerOptions{Repository: mockRepo})

	ctx := context.Background()
	userID := uint(1)
	swipedID := uint(2)

	mockRepo.EXPECT().GetUserByID(ctx, int(userID)).Return(&model.User{}, nil)

	mockRepo.EXPECT().IsUserPremium(ctx, userID).Return(true, nil)

	mockRepo.EXPECT().GetUserByID(ctx, int(swipedID)).Return(&model.User{}, nil)

	expectedSwipe := &model.Swipe{
		SwiperID: userID,
		SwipedID: swipedID,
		IsLike:   true,
	}
	mockRepo.EXPECT().CreateSwipe(ctx, expectedSwipe).Return(nil)

	mockRepo.EXPECT().CheckMatch(ctx, userID, swipedID).Return(true, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("userID", userID)

	c.Request = httptest.NewRequest("POST", "/swipe", strings.NewReader(`{"user_id": 2, "is_like": true}`))

	server.Swipe(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "It's a match!", response["message"])
	assert.Equal(t, true, response["is_match"])
}

// Non premium user successfully swipes right and NOT match
func TestPremiumUserSwipeRightNotMatch(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockRepositoryInterface(ctrl)
	server := handler.NewServer(handler.NewServerOptions{Repository: mockRepo})

	ctx := context.Background()
	userID := uint(1)
	swipedID := uint(2)

	mockRepo.EXPECT().GetUserByID(ctx, int(userID)).Return(&model.User{}, nil)

	mockRepo.EXPECT().IsUserPremium(ctx, userID).Return(true, nil)

	mockRepo.EXPECT().GetUserByID(ctx, int(swipedID)).Return(&model.User{}, nil)

	expectedSwipe := &model.Swipe{
		SwiperID: userID,
		SwipedID: swipedID,
		IsLike:   true,
	}
	mockRepo.EXPECT().CreateSwipe(ctx, expectedSwipe).Return(nil)

	mockRepo.EXPECT().CheckMatch(ctx, userID, swipedID).Return(false, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("userID", userID)

	c.Request = httptest.NewRequest("POST", "/swipe", strings.NewReader(`{"user_id": 2, "is_like": true}`))

	server.Swipe(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Swipe recorded successfully", response["message"])
	assert.Equal(t, false, response["is_match"])
}

// Premium user successfully swipes right return to many requests
func TestNonPremiumUserSwipeRightTooManyRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockRepositoryInterface(ctrl)
	server := handler.NewServer(handler.NewServerOptions{Repository: mockRepo})

	ctx := context.Background()
	userID := uint(1)

	mockRepo.EXPECT().GetUserByID(ctx, int(userID)).Return(&model.User{}, nil)

	mockRepo.EXPECT().IsUserPremium(ctx, userID).Return(false, nil)

	mockRepo.EXPECT().GetDailySwipeCount(ctx, userID).Return(10, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("userID", userID)

	c.Request = httptest.NewRequest("POST", "/swipe", strings.NewReader(`{"user_id": 2, "is_like": true}`))

	server.Swipe(c)

	assert.Equal(t, http.StatusTooManyRequests, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Upgrade to premium to remove the swipe limit!", response["message"])
}

func TestUserNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockRepositoryInterface(ctrl)
	server := handler.NewServer(handler.NewServerOptions{Repository: mockRepo})

	ctx := context.Background()
	userID := uint(1)

	mockRepo.EXPECT().GetUserByID(ctx, int(userID)).Return(nil, errors.New("not found"))

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("userID", userID)

	c.Request = httptest.NewRequest("POST", "/swipe", strings.NewReader(`{"user_id": 2, "is_like": true}`))

	server.Swipe(c)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
}

func TestFailedGetPremiumStatus(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockRepositoryInterface(ctrl)
	server := handler.NewServer(handler.NewServerOptions{Repository: mockRepo})

	ctx := context.Background()
	userID := uint(1)

	mockRepo.EXPECT().GetUserByID(ctx, int(userID)).Return(&model.User{}, nil)

	mockRepo.EXPECT().IsUserPremium(ctx, userID).Return(false, errors.New("failed"))

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("userID", userID)

	c.Request = httptest.NewRequest("POST", "/swipe", strings.NewReader(`{"user_id": 2, "is_like": true}`))

	server.Swipe(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
}
