package services

import (
	"patient_portal/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

type AuthService struct {
    db        *gorm.DB
    jwtSecret string
}

func NewAuthService(db *gorm.DB, secret string) *AuthService {
    return &AuthService{db, secret}
}

func (s *AuthService) Authenticate(username, password string) (string, error) {
    var user models.User
    if err := s.db.Where("username = ?", username).First(&user).Error; err != nil {
        return "", err
    }
    if user.Password != password {
        return "", gorm.ErrRecordNotFound
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "sub":  user.ID,
        "role": user.Role,
        "exp":  time.Now().Add(72 * time.Hour).Unix(),
    })
    return token.SignedString([]byte(s.jwtSecret))
}
