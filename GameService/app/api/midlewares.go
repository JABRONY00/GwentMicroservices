package api

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitMiddlewares(r *gin.Engine, dbPool *pgxpool.Pool) {
	r.Use(func(c *gin.Context) {
		c.Set("dbPool", dbPool)
		c.Next()
	})
	r.Use(func(c *gin.Context) {
		startTime := time.Now()
		c.Set("startTime", startTime)
		c.Next()
	})
}
