package api

import (
	"GwentMicroservices/UserService/app/api/handlers"
	"GwentMicroservices/UserService/app/api/query"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Routes(r *gin.Engine, db *pgxpool.Pool) {
	InitMiddlewares(r, db)
	query.TransferDB(db)
	r.POST("/user/sign-up", handlers.UserSignUp)
	r.GET("/user/login", handlers.UserLogin)
	r.GET("/user/logout", handlers.UserLogout)
}
