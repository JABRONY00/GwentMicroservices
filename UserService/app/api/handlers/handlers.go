package handlers

import (
	"GwentMicroservices/UserService/app/api/models"
	"GwentMicroservices/UserService/app/api/services"
	"GwentMicroservices/UserService/app/helpers"
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

	exists, err := services.PlayerExistanceCheck(player.Name)
	switch {
	case err != nil:
		{
			log.HttpLog(c, log.Error, http.StatusInternalServerError, err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}
	case exists:
		{
			log.HttpLog(c, log.Warn, http.StatusBadRequest, "invalid name")
			c.JSON(http.StatusBadRequest, gin.H{"error": "player name is already reserved"})
			return
		}
	case helpers.ValidateEmail(player.Email):
		{
			log.HttpLog(c, log.Warn, http.StatusBadRequest, "Invalid email")
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid email"})
			return
		}
	case helpers.ValidatePassword(player.Password):
		{
			log.HttpLog(c, log.Warn, http.StatusBadRequest, "invalid password")
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
			return
		}
	}

	err = services.CreatePlayer(&player)
	if err != nil {
		log.HttpLog(c, log.Error, http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	log.HttpLog(c, log.Info, http.StatusOK, "new user created")
	c.JSON(http.StatusOK, gin.H{"message": "registrated sucsessfully"})

	signedString, limit, err := services.CreateToken(player.ID)
	if err != nil {
		log.HttpLog(c, log.Info, http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	c.SetCookie("token", signedString, limit, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "logged in"})
}

func UserLogin(c *gin.Context) {
	var player models.PlayerInfoPassword
	err := c.ShouldBindJSON(&player)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = services.AuthPlayer(&player)
	switch {
	case err.Error() == "bad password":
		{
			c.JSON(http.StatusBadRequest, gin.H{"error": "bad password"})
			return
		}
	case err != nil:
		{
			log.HttpLog(c, log.Info, http.StatusInternalServerError, err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}
	}

	signedString, limit, err := services.CreateToken(player.ID)
	if err != nil {
		log.HttpLog(c, log.Info, http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.SetCookie("token", signedString, limit, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "logged in"})
}

func UserLogout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "user logged out"})
}
