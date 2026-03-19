package services

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"openim/internal/middleware"
)

func TestUserRegister_BadRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(middleware.RequestIDMiddleware())
	r.POST("/api/register", UserRegister)

	tests := []struct {
		name   string
		body   string
		expect int
	}{
		{"empty body", "", http.StatusBadRequest},
		{"invalid json", "{", http.StatusBadRequest},
		{"missing user_name", `{"password":"123"}`, http.StatusBadRequest},
		{"missing password", `{"user_name":"testuser"}`, http.StatusBadRequest},
		{"empty user_name", `{"user_name":"","password":"123"}`, http.StatusBadRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/api/register", strings.NewReader(tt.body))
			req.Header.Set("Content-Type", "application/json")
			r.ServeHTTP(w, req)
			if w.Code != tt.expect {
				t.Errorf("code = %d, want %d, body: %s", w.Code, tt.expect, w.Body.String())
			}
		})
	}
}
