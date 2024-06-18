package connections

import (
	"GwentMicroservices/GameService/app/engine"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func MatchMaker(dbPool *pgxpool.Pool) {
	var pair []WaitingClient
	//var disconnected []string
	for {
		WaitingClients.RWMutex.Lock()
		for key, value := range WaitingClients.Content {
			if value != nil && len(value.(chan bool)) != 1 {
				pair = append(pair, WaitingClient{Name: key, Ch: value.(chan bool)})
			}
			if len(pair) == 2 {
				break
			}
		}
		WaitingClients.RWMutex.Unlock()
		for index := range pair {
			WaitingClients.Delete(pair[index].Name)
		}
		NewTable(dbPool, pair[0].Name, pair[1].Name)
		pair[0].Ch <- true
		pair[1].Ch <- true
		close(pair[0].Ch)
		close(pair[1].Ch)
		pair = nil
		time.Sleep(time.Second * 1)
	}
}

func NewTable(dbPool *pgxpool.Pool, clientname1 string, clientname2 string) {
	client1 := ActiveClients.Get(clientname1).(Client)
	client2 := ActiveClients.Get(clientname2).(Client)
	newTable := engine.NewTable(dbPool,
		engine.Client{Name: client1.Name, Conn: client1.Conn},
		engine.Client{Name: client2.Name, Conn: client2.Conn})
	client1.TableID = newTable.TableID
	client2.TableID = newTable.TableID
	ActiveClients.Set(clientname1, client1)
	ActiveClients.Set(clientname2, client2)
	fmt.Println("New Table created!")
}
