package server

import (
	"be-tasking/constanta"
	"be-tasking/helper"
	"net/http"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// Middleware function for origin permit
func CORSMiddleware(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT, DELETE")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}

	c.Next()
}

// Middleware function for token validation
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")

		if govalidator.IsNull(authorizationHeader) {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Missing authorization header"})
			c.Abort()
			return
		}

		// Check if the Authorization header starts with "Bearer "
		if !strings.HasPrefix(authorizationHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid authorization format"})
			c.Abort()
			return
		}

		// Extract the token part from the Authorization header
		tokenString := authorizationHeader[len("Bearer "):]

		if err := helper.ValidateToken(tokenString); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			c.Abort()
			return
		}

		c.Next()
	}
}

func PelaksanaOnlyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")
		tokenString := authorizationHeader[len("Bearer "):]
		claim, _ := helper.GetTokenClaims(tokenString)

		if claim.Role != constanta.RoleTypePelaksana {
			c.AbortWithStatusJSON(403, gin.H{"error": "Access denied for this resource."})
			return
		}

		c.Next()
	}
}

func LeaderOnlyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")
		tokenString := authorizationHeader[len("Bearer "):]
		claim, _ := helper.GetTokenClaims(tokenString)

		if claim.Role != constanta.RoleTypeLeader {
			c.AbortWithStatusJSON(403, gin.H{"error": "Access denied for this resource."})
			return
		}

		c.Next()
	}
}

func LeaderAndPelaksanaOnlyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")
		tokenString := authorizationHeader[len("Bearer "):]
		claim, _ := helper.GetTokenClaims(tokenString)

		if claim.Role != constanta.RoleTypeLeader && claim.Role != constanta.RoleTypePelaksana {
			c.AbortWithStatusJSON(403, gin.H{"error": "Access denied for this resource."})
			return
		}

		c.Next()
	}
}
