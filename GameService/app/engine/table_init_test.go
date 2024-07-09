package engine_test

import (
	"GwentMicroservices/GameService/app/api/models"
	"GwentMicroservices/GameService/app/engine"
	"testing"
)

func TestNewTable(t *testing.T) {

	t.Logf("New table created! Info: %+v", engine.NewTable(
		&models.Client{Name: "name1"},
		&models.Client{Name: "name2"},
	),
	)
}
