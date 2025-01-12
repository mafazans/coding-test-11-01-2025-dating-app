package handler_test

// func TestRegisterValidRequest(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockRepo := repository.NewMockRepositoryInterface(ctrl)
// 	handler := handler.NewServer(handler.NewServerOptions{Repository: mockRepo})

// 	mockRepo.EXPECT().
// 		CreateUser(gomock.Any()).
// 		Return(nil)

// 	w := httptest.NewRecorder()
// 	c, _ := gin.CreateTestContext(w)

// 	reqBody := `{"name": "test", "password": "password123"}`
// 	c.Request, _ = http.NewRequest(
// 		http.MethodPost,
// 		"/register",
// 		bytes.NewBufferString(reqBody),
// 	)
// 	c.Request.Header.Set("Content-Type", "application/json")

// 	handler.Register(c)

// 	assert.Equal(t, http.StatusCreated, w.Code)
// 	assert.Contains(t, w.Body.String(), " registered successfully")
// }
