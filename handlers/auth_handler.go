package handlers

import (
	"net/http"
	"patient_portal/services"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
    AuthSvc *services.AuthService
}

func NewAuthHandler(s *services.AuthService) *AuthHandler {
    return &AuthHandler{s}
}

func (h *AuthHandler) Login(c *gin.Context) {
    var req struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    token, err := h.AuthSvc.Authenticate(req.Username, req.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"token": token})
}
