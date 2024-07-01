package initializers

import (
	"GwentMicroservices/GameService/app/api/services"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitGameServer(dbPool *pgxpool.Pool) {
	services.ActiveGameTables.Init()
	services.ActiveClients.Init()
	services.WaitingClients.Init()
	go services.MatchMaker(dbPool)
}
