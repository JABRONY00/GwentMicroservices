package services

import (
	"GwentMicroservices/GameService/app/api/query"
	"GwentMicroservices/GameService/app/engine"
)

func NewTable(clientname1 string, clientname2 string) error {
	go GameConnection(clientname1)
	go GameConnection(clientname2)

	client1 := ActiveClients.Get(clientname1)
	client2 := ActiveClients.Get(clientname2)

	newTable := engine.NewTable(
		&client1,
		&client2,
	)

	presets, err := query.GetPlayersPreset(clientname1, clientname2)
	if err != nil {
		return err
	}

	newTable.InitGame(presets)

	ActiveGameTables.Set(newTable.TableID, newTable)

	client1.TableID = newTable.TableID
	client2.TableID = newTable.TableID

	ActiveClients.Set(clientname1, client1)
	ActiveClients.Set(clientname2, client2)

	return nil
}
