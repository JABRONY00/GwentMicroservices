package query

import (
	"GwentMicroservices/GameService/app/api/models"
	"context"
)

func GetPlayerNameByID(id string) (string, error) {

	rows := DB.QueryRow(context.Background(), "SELECT name FROM players WHERE id = $1", id)
	var name string
	err := rows.Scan(&name)
	return name, err
}

func GetPlayersPreset(name1 string, name2 string) (map[string]models.PlayerPreset, error) {

	rows, err := DB.Query(context.Background(), "SELECT name, race, stack FROM presets WHERE name IN('$1', '$2')", name1, name2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	presets := make(map[string]models.PlayerPreset)
	for rows.Next() {
		var name string
		var preset models.PlayerPreset
		err := rows.Scan(
			&name,
			&preset.Race,
			&preset.Stack,
		)
		if err != nil {
			return nil, err
		}

		presets[name] = preset
	}

	return presets, nil

}
