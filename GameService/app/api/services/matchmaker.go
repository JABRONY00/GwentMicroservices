package services

import (
	"GwentMicroservices/GameService/app/api/models"
	"GwentMicroservices/GameService/app/engine"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func MatchMaker(dbPool *pgxpool.Pool) {
	var pair []models.WaitingClient
	var disconnected []string
	for {
		WaitingClients.RWMutex.Lock()
		for key, value := range WaitingClients.Content {

			if value != nil {
				pair = append(pair, models.WaitingClient{Name: key, Ch: value.(chan bool)})
			} else {
				disconnected = append(disconnected, key)
			}
			if len(pair) == 2 {
				break
			}
		}
		WaitingClients.RWMutex.Unlock()
		for index := range pair {
			pair[index].Ch <- true
			WaitingClients.Delete(pair[index].Name)
		}
		NewTable(dbPool, pair[0].Name, pair[1].Name)
		if len(disconnected) != 0 {
			for index := range disconnected {
				WaitingClients.Delete(disconnected[index])
			}
		}
		pair = nil
		time.Sleep(time.Second * 1)
	}
}

func NewTable(dbPool *pgxpool.Pool, clientname1 string, clientname2 string) {
	go ReadConnection(clientname1)
	go ReadConnection(clientname2)
	client1 := ActiveClients.Get(clientname1).(models.Client)
	client2 := ActiveClients.Get(clientname2).(models.Client)
	newTable := engine.NewTable(dbPool,
		engine.Client{Name: client1.Name, Conn: client1.Conn},
		engine.Client{Name: client2.Name, Conn: client2.Conn})
	client1.TableID = newTable.TableID
	client2.TableID = newTable.TableID
	ActiveClients.Set(clientname1, client1)
	ActiveClients.Set(clientname2, client2)

}
