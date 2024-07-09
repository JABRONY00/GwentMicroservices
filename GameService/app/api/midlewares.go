package api

import (
	"GwentMicroservices/GameService/app/api/models"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitMiddlewares(r *gin.Engine, dbPool *pgxpool.Pool) {

	r.Use(func(c *gin.Context) {
		startTime := time.Now()
		c.Set("startTime", startTime)
		c.Next()
	})
}

func AuthMidlewares() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		token, err := jwt.ParseWithClaims(cookie, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("gwent"), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}
		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}
		claims := token.Claims.(*models.Claims)
		c.Set("player", claims.Subject)
		c.Next()
	}
}
