package query

import (
	"GwentMicroservices/UserService/app/api/models"
	"context"
)

func CheckPlayerExists(name string) (bool, error) {
	res := DB.QueryRow(context.Background(), "SELECT DISTINCT name FROM players WHERE name = $1", name)
	err := res.Scan()
	switch {
	case err == nil:
		{
			return true, nil
		}
	case err.Error() == "no rows in result set":
		{
			return false, nil
		}
	}

	return false, err
}

func InsertPlayer(player *models.PlayerInfoPassword) error {

	_, err := DB.Exec(context.Background(), "INSERT INTO players VALUES ($1, $2, $3, $4)", player.ID, player.Name, player.Email, player.PasswordHash)

	return err
}

func GetPlayerForAuth(player *models.PlayerInfoPassword) error {
	rows := DB.QueryRow(context.Background(), "SELECT id, password FROM players WHERE email = $1", player.Email)
	err := rows.Scan(&player.ID, &player.PasswordHash)

	return err
}
