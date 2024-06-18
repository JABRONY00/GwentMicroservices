package api

import (
	"GwentMicroservices/GameService/app/api/connections"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Routes(r *gin.Engine, dbPool *pgxpool.Pool) {
	InitMiddlewares(r, dbPool)
	r.POST("/game/start", connections.NewConnection)

}
