package helpers

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

func GenerateToken(id uint, email string) (string, error) {
	jwtSecret := viper.GetString("jwt.secret")

	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := parseToken.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", errors.New("failed to sign token")
	}

	return signed, nil
}

func VerifyToken(c *gin.Context) (interface{}, error) {
	jwtSecret := viper.GetString("jwt.secret")

	errResponse := errors.New("token has invalid")
	headerToken := c.Request.Header.Get("authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")

	if !bearer {
		return nil, errResponse
	}

	jwtString := strings.Split(headerToken, " ")[1]

	token, err := jwt.Parse(jwtString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, errResponse
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errResponse
	}

	return token.Claims.(jwt.MapClaims), nil
}
