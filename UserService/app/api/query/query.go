package query

import (
	"GwentMicroservices/UserService/app/api/models"
	"context"
)

func CheckPlayerExists(name string) error {
	res := DB.QueryRow(context.Background(), "SELECT DISTINCT name FROM players WHERE name = $1", name)
	var existingName string
	err := res.Scan(&existingName)
	return err
}

func InsertPlayer(player *models.PlayerInfoPassword) error {

	_, err := DB.Exec(context.Background(), "INSERT INTO players VALUES ($1, $2, $3, $4)", player.ID, player.Name, player.Email, player.PasswordHash)

	return err
}

func GetPlayerForAuth(player *models.PlayerInfoPassword) error {
	rows := DB.QueryRow(context.Background(), "SELECT id, password_hash FROM players WHERE email = $1", player.Email)
	err := rows.Scan(&player.ID, &player.PasswordHash)

	return err
}
