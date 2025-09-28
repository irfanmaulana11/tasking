package helper

import (
	"errors"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/dgrijalva/jwt-go"
)

type TokenClaims struct {
	Sub         int    `json:"sub"`
	EmployeeID  string `json:"employee_id"`
	UserName    string `json:"username"`
	DisplayName string `json:"display_name"`
	Exp         int64  `json:"exp"`
	Role        string `json:"role"`
	jwt.StandardClaims
}

var jwtKey []byte

func InitJWTKey(key []byte) {
	jwtKey = key
}

func CreateToken(claim TokenClaims) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = claim.UserName
	claims["display_name"] = claim.DisplayName
	claims["employee_id"] = claim.EmployeeID
	claims["role"] = claim.Role
	claims["sub"] = claim.Sub
	claims["exp"] = claim.Exp

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return errors.New("Invalid Token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("Invalid Token")
	}

	// Check the expiration time
	expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)
	if time.Now().After(expirationTime) {
		return errors.New("Token is expired")
	}

	return nil
}

// get token from header
func GetTokenHeader(c *gin.Context) string {
	authorizationHeader := c.GetHeader("Authorization")
	tokenString := authorizationHeader[len("Bearer "):]
	return tokenString
}

// access token claims
func GetTokenClaims(tokenString string) (*TokenClaims, error) {
	claim := &TokenClaims{}

	// Parse the token into the custom Claims struct
	_, err := jwt.ParseWithClaims(tokenString, claim, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	return claim, err
}
