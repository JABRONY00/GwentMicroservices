package services

import (
	"GwentMicroservices/GameService/app/api/models"
	"GwentMicroservices/GameService/app/api/query"
	"GwentMicroservices/GameService/app/engine"
	"GwentMicroservices/GameService/app/helpers/log"
	"encoding/json"
	"fmt"
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func NewConnection(c *gin.Context) {
	if uint(len(ActiveGameTables.Content)) >= TablesLimit {
		log.HttpLog(c, log.Error, http.StatusTooManyRequests, "Tables limit reached! Connection declined!")
		c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many players"})
		return
	}
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	connection, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.HttpLog(c, log.Error, http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	playerID, ok := c.Keys["ID"].(string)
	if !ok {
		log.HttpLog(c, log.Error, http.StatusInternalServerError, "Failed to get player ID")
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	name, err := query.GetPlayerNameByID(playerID)
	if err != nil {
		log.HttpLog(c, log.Error, http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	ActiveClients.Set(name, models.Client{
		Name:    name,
		Conn:    &models.Connection{Conn: connection},
		TableID: "",
	})
	ch := make(chan struct{})
	go WaitingConnection(name, ch)
	WaitingClients.Set(name, ch)
}

func WaitingConnection(name string, matchmakerChan chan struct{}) {
	client := ActiveClients.Get(name)
	client.SendJson("Wait for second player")

	for {
		_, isOpen := <-matchmakerChan
		switch {
		case !isOpen:
			{
				return
			}
		default:
			{
				if _, _, err := client.Conn.ReadMessage(); err != nil {
					matchmakerChan <- struct{}{}
					return
				}
			}
		}
	}
}

func GameConnection(player string) {
	client := ActiveClients.Get(player)
	defer client.Conn.Close()
	client.SendJson(models.ResponseData{Instr: "Game is running..."})
	for {
		_, message, err := client.Conn.ReadMessage()
		if err != nil {
			fmt.Println("Failed to read message: ", err)
			break
		}
		var reqBody engine.RequestData
		err = json.Unmarshal(message, &reqBody)
		if err != nil {
			fmt.Println("Unmarshal Error: ", err)
			continue
		}
		if reqBody.Instr != "" {
			/////////////// Move processing /////////////////////
			t := ActiveGameTables.Get(client.TableID)

			// Move validation
			switch {
			case reqBody.Instr == "check":
				{
					client.SendJson(models.ResponseData{Instr: engine.Instr.Refresh, Data: t.GetTableInfo(client.Name)})
					continue
				}
			case reqBody.Instr == "check2":
				{

					continue
				}
			case reqBody.Instr == engine.Instr.Meta:
				{
					t.MetaResponse(client.Name, reqBody.CardID)
					continue
				}
			case t.Pm.ActPlr != client.Name:
				fallthrough
			case t.Pm.Instr != engine.Instr.Move &&
				(reqBody.Instr == engine.Instr.PutCard ||
					reqBody.Instr == engine.Instr.LBonus ||
					reqBody.Instr == engine.Instr.Pass):
				fallthrough
			case !slices.Contains(t.Pm.IDs, reqBody.CardID) && reqBody.CardID != 0:
				{
					client.SendJson(models.ResponseData{Instr: engine.Instr.ForbMove})
					continue
				}
			}
			// Move validation

			err := t.MoveRouter(reqBody)
			if err != nil {
				t.Players[t.Pm.ActPlr].SendJson(models.ResponseData{Instr: engine.Instr.ForbMove})
				continue
			}

			ActiveGameTables.Set(t.TableID, t)
			t.RefreshTable()
			t.Players[t.Pm.ActPlr].SendJson(models.ResponseData{Instr: engine.Instr.Move, Data: t.Pm.IDs})

			switch {
			case t.Players[t.Pm.PasPlr].PassFlag:
				{
					t.Players[t.Pm.PasPlr].SendJson(models.ResponseData{Instr: engine.Instr.Pass})
				}
			case t.Pm.Instr == engine.Instr.Move:
				{
					t.Players[t.Pm.PasPlr].SendJson(models.ResponseData{Instr: engine.Instr.Wait})
				}
			}
			/////////////// Move processing //////////////
		}
	}
}
