package jwt_token

import (
	"crm-platform-management-api/internal/models"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func GenerateToken(user *models.CRMUser) *jwt.Token {
	//  generates token for auth login user
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":         user.ID,
		"expiration": time.Now().Add(time.Hour * 24).Unix(),
	})

	return claims
}
