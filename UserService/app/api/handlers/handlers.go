package handlers

import (
	"GwentMicroservices/UserService/app/helpers"
	"GwentMicroservices/UserService/app/helpers/log"
	"context"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

func UserRegistration(c *gin.Context) {
	var info UserCredentials
	err := c.ShouldBindJSON(&info)
	if err != nil {
		log.HttpLog(c, log.Warn, http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Body"})
		return
	}
	dbPool := c.MustGet("dbPool").(*pgxpool.Pool)
	switch {
	case !helpers.ValidateName(info.Name, dbPool):
		fallthrough
	case !helpers.ValidatePassword(info.Password):
		{
			log.HttpLog(c, log.Warn, http.StatusBadRequest, "Invalid user name or password")
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user name or password"})
			return
		}
	}
	info.HashedPassword, err = bcrypt.GenerateFromPassword([]byte(info.Password), 14)
	if err != nil {
		log.HttpLog(c, log.Error, http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	_, err = dbPool.Exec(context.Background(), "INSERT INTO players VALUES ($1, $2)", info.Name, info.HashedPassword)
	if err != nil {
		log.HttpLog(c, log.Error, http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	log.HttpLog(c, log.Info, http.StatusOK, "new user registrated")
	c.JSON(http.StatusOK, gin.H{"message": "registrated sucsessfully"})
	claims := &Claims{
		StandardClaims: jwt.StandardClaims{
			Subject:   info.Name,
			ExpiresAt: time.Now().Add(2 * time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("gwent"))
	if err != nil {
		log.HttpLog(c, log.Info, http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	c.SetCookie("token", tokenString, int(claims.StandardClaims.ExpiresAt), "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "logged in"})
}

func UserLogin(c *gin.Context) {
	var userInfo UserCredentials
	err := c.ShouldBindJSON(&userInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dbpool, _ := c.Get("dbpool")
	db := dbpool.(*pgxpool.Pool)
	info := db.QueryRow(context.Background(), "SELECT password FROM players WHERE players_name = $1", userInfo.Name)
	err = info.Scan(&userInfo.HashedPassword)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "invalid user name"})
		return
	}
	err = bcrypt.CompareHashAndPassword(userInfo.HashedPassword, []byte(userInfo.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
		return
	}
	claims := &Claims{
		StandardClaims: jwt.StandardClaims{
			Subject:   userInfo.Name,
			ExpiresAt: time.Now().Add(2 * time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("gwent"))
	if err != nil {
		log.HttpLog(c, log.Info, http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	c.SetCookie("token", tokenString, int(claims.StandardClaims.ExpiresAt), "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "logged in"})
}

func UserLogout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "user logged out"})
}
