package api

import (
	"GwentMicroservices/UserService/app/api/handlers"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Routes(r *gin.Engine, dbPool *pgxpool.Pool) {
	InitMiddlewares(r, dbPool)
	r.POST("/user/registration", handlers.UserRegistration)
	r.GET("/user/login", handlers.UserLogin)
	r.GET("/user/logout", handlers.UserLogout)
}
