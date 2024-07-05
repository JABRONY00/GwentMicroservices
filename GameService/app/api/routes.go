package api

import (
	"GwentMicroservices/GameService/app/api/query"
	"GwentMicroservices/GameService/app/api/services"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Routes(r *gin.Engine, dbPool *pgxpool.Pool) {
	InitMiddlewares(r, dbPool)
	query.TransferDB(dbPool)
	r.POST("/game/start", services.NewConnection)

}
