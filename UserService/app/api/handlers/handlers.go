package handlers

import (
	"GwentMicroservices/UserService/app/api/models"
	"GwentMicroservices/UserService/app/api/services"
	log "GwentMicroservices/UserService/app/helpers/log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserSignUp(c *gin.Context) {
	var player models.PlayerInfoPassword
	err := c.ShouldBindJSON(&player)
	if err != nil {
		log.HttpLog(c, log.Warn, http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Body"})
		return
	}

	err = services.RegistrationInfoValidation(player)
	if err != nil {
		log.HttpLog(c, log.Warn, http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = services.PlayerExistanceCheck(player.Name)
	switch {
	case err == nil:
		{
			log.HttpLog(c, log.Warn, http.StatusBadRequest, "invalid name")
			c.JSON(http.StatusBadRequest, gin.H{"error": "player name is already reserved"})
			return
		}
	case err.Error() != "no rows in result set":
		{
			log.HttpLog(c, log.Error, http.StatusInternalServerError, err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{})
			return
		}
	}

	err = services.CreatePlayer(&player)
	if err != nil {
		log.HttpLog(c, log.Error, http.StatusInternalServerError, err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	log.HttpLog(c, log.Info, http.StatusOK, "new user created")
	c.JSON(http.StatusOK, gin.H{"message": "registrated sucsessfully"})

	signedString, limit, err := services.CreateToken(player.ID)
	if err != nil {
		log.HttpLog(c, log.Info, http.StatusInternalServerError, err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.SetCookie("token", signedString, limit, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "logged in"})
}

func UserLogin(c *gin.Context) {
	var player models.PlayerInfoPassword
	err := c.ShouldBindJSON(&player)
	switch {
	case err != nil:
		fallthrough
	case player.Email == "":
		fallthrough
	case player.Password == "":
		{
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
			return
		}
	}

	err = services.AuthPlayer(&player)
	switch {
	case err == nil:
		{
			break
		}
	case err.Error() == "no rows in result set":
		{
			c.JSON(http.StatusBadRequest, gin.H{"error": "player does not exist"})
			return
		}
	case err.Error() == "bad password":
		{
			c.JSON(http.StatusBadRequest, gin.H{"error": "bad password"})
			return
		}
	default:
		{
			log.HttpLog(c, log.Info, http.StatusInternalServerError, err.Error())
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}

	signedString, limit, err := services.CreateToken(player.ID)
	if err != nil {
		log.HttpLog(c, log.Info, http.StatusInternalServerError, err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.SetCookie("token", signedString, limit, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "logged in"})
}

func UserLogout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "user logged out"})
}
