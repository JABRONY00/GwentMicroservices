package services

import (
	"GwentMicroservices/UserService/app/api/models"
	"GwentMicroservices/UserService/app/api/query"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func PlayerExistanceCheck(name string) (bool, error) {
	exists, err := query.CheckPlayerExists(name)

	return exists, err
}

func CreatePlayer(player *models.PlayerInfoPassword) error {
	var err error
	player.PasswordHash, err = bcrypt.GenerateFromPassword([]byte(player.Password), 10)
	if err != nil {
		return err
	}

	player.ID = uuid.New().String()
	if player.ID == "" {
		return errors.New("failed to generate new uuid")
	}

	err = query.InsertPlayer(player)

	return err
}

func CreateToken(id string) (string, int, error) {

	claims := &models.Claims{
		StandardClaims: jwt.StandardClaims{
			Subject:   id,
			ExpiresAt: time.Now().Add(2 * time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString([]byte("gwent"))

	return signedString, int(claims.ExpiresAt), err
}

func AuthPlayer(player *models.PlayerInfoPassword) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(player.Password), 10)
	if err != nil {
		return err
	}

	err = query.GetPlayerForAuth(player)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword(player.PasswordHash, passwordHash)
	if err != nil {
		return errors.New("bad password")
	}

	return nil
}
