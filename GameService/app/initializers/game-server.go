package initializers

import (
	"GwentMicroservices/GameService/app/api/connections"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitGameServer(dbPool *pgxpool.Pool) {
	connections.ActiveGameTables.Init()
	connections.ActiveClients.Init()
	connections.WaitingClients.Init()
	go connections.MatchMaker(dbPool)
}
