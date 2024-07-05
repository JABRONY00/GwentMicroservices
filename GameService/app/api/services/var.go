package services

import (
	"GwentMicroservices/GameService/app/api/models"
	"GwentMicroservices/GameService/app/engine"
)

const TablesLimit uint = 5

var ActiveClients models.ConcMap[models.Client]
var WaitingClients models.ConcMap[chan struct{}]
var ActiveGameTables models.ConcMap[*engine.Table]
