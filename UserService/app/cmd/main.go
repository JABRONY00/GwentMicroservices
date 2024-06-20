package main

import (
	"GwentMicroservices/UserService/app/api"
	"GwentMicroservices/UserService/app/helpers"
	"GwentMicroservices/UserService/app/initializers"
	"fmt"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var (
	SERVER_PORT = helpers.GetEnv("SERVER_PORT")
	SERVER_HOST = helpers.GetEnv("SERVER_HOST")
)

func init() {
	helpers.CheckRequiredEnvs()
	initializers.InitLogger()
}

func main() {
	db := initializers.DbConnection()
	defer db.Close()
	router := gin.Default()
	api.Routes(router, db)
	err := router.Run(fmt.Sprintf("%v:%v", SERVER_HOST, SERVER_PORT))
	if err != nil {
		log.Panicf("Server listen err: %v", err)
	}
	log.Infof("Server started succsessfully")
}
