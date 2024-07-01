package services

import "GwentMicroservices/GameService/app/api/models"

const TablesLimit uint = 5

var ActiveClients models.ConcMap
var WaitingClients models.ConcMap
var ActiveGameTables models.ConcMap
