package services

import (
	"GwentMicroservices/GameService/app/api/models"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func MatchMaker(dbPool *pgxpool.Pool) {
generalLoop:
	for {
		var disconnected []string
		pair := make([]models.WaitingClient, 0, 2)
		WaitingClients.RWMutex.RLock()
		if len(WaitingClients.Content) < 2 {
			WaitingClients.RWMutex.RUnlock()
			time.Sleep(time.Second * 1)
			continue generalLoop
		}

	innerLoop:
		for key, value := range WaitingClients.Content {
			switch {
			case len(pair) == 2:
				break innerLoop
			case len(value) != 0:
				{
					close(value)
					disconnected = append(disconnected, key)
				}
			default:
				{
					pair = append(pair, models.WaitingClient{Name: key, Ch: value})
				}
			}
		}
		WaitingClients.RWMutex.RUnlock()

		if len(pair[0].Ch) == 0 && len(pair[1].Ch) == 0 {
			close(pair[0].Ch)
			close(pair[1].Ch)
			NewTable(pair[0].Name, pair[1].Name)
			WaitingClients.Delete(pair[0].Name)
			WaitingClients.Delete(pair[1].Name)
		} else {
			if len(pair[0].Ch) != 0 {
				close(pair[0].Ch)
				disconnected = append(disconnected, pair[0].Name)
			}
			if len(pair[1].Ch) != 0 {
				close(pair[1].Ch)
				disconnected = append(disconnected, pair[1].Name)
			}
		}

		if len(disconnected) != 0 {
			WaitingClients.RWMutex.Lock()
			for index := range disconnected {
				delete(WaitingClients.Content, disconnected[index])
			}
			WaitingClients.RWMutex.Unlock()
		}

		pair = nil
	}
}
