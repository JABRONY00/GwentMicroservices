package main

import (
	"GwentMicroservices/GameService/app/api"
	"GwentMicroservices/GameService/app/initializers"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	initializers.InitLogger()
	dbPool := initializers.DbConnection()
	defer dbPool.Close()
	initializers.InitGameServer(dbPool)
	router := gin.Default()
	api.Routes(router, dbPool)
	err := router.Run("localhost:4001")
	if err != nil {
		log.Panicf("Server listen err: %v", err)
	}
	log.Infof("Server started on localhost:4000")
}
