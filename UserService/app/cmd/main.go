package main

import (
	"GwentMicroservices/UserService/app/api"
	"GwentMicroservices/UserService/app/initializers"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func init() {
	initializers.InitLogger()
}

func main() {
	dbPool := initializers.DbConnection()
	defer dbPool.Close()
	router := gin.Default()
	api.Routes(router, dbPool)
	err := router.Run("localhost:4000")
	if err != nil {
		log.Panicf("Server listen err: %v", err)
	}
	log.Infof("Server started on localhost:4000")
}
