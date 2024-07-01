package services

import (
	"GwentMicroservices/GameService/app/api/models"
	"GwentMicroservices/GameService/app/engine"
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func NewConnection(c *gin.Context) {
	if uint(len(ActiveGameTables.Content)) >= TablesLimit {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Too many players"})
		fmt.Println("Tables limit reached! Connection declined!")
		return
	}
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	connection, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}
	name := c.MustGet("player").(string)
	ActiveClients.Set(name, models.Client{
		Name:    name,
		Conn:    connection,
		TableID: "",
	})
	ch := make(chan bool)
	go WaitingConnection(name, ch)
	WaitingClients.Set(name, ch)
}

func WaitingConnection(name string, matchmakerChan chan bool) {
	startTime := time.Now()
	ticker := time.NewTicker(time.Second)
	stopTimer := make(chan bool)
	client := ActiveClients.Get(name).(engine.Client)
	client.Conn.Mut.Lock()
	client.Conn.WriteJSON("Wait for second player")
	client.Conn.Mut.Unlock()
	go func() {
		for {
			select {
			case <-stopTimer:
				return
			case t := <-ticker.C:
				{
					client.Conn.Mut.Lock()
					client.Conn.WriteJSON(engine.ResponseData{Instr: "waiting-time", Data: t.Sub(startTime).Round(time.Second).String()})
					client.Conn.Mut.Unlock()
				}
			}
		}
	}()
	for {
		select {
		case <-matchmakerChan:
			{
				close(matchmakerChan)
				matchmakerChan = nil
				stopTimer <- true
				ticker.Stop()
				return
			}
		default:
			{
				_, _, err := client.Conn.ReadMessage()
				if err != nil {
					close(matchmakerChan)
					matchmakerChan = nil
					stopTimer <- true
					client.Conn.Close()
					close(stopTimer)
					ticker.Stop()
					return
				}
			}
		}
	}
}

func ReadConnection(player string) {
	client := ActiveClients.Get(player).(models.Client)
	defer client.Conn.Close()
	client.Conn.WriteJSON(engine.ResponseData{Instr: "Game is running..."})
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
			t := ActiveGameTables.Get(client.TableID).(*engine.Table)
			// Move validation
			switch {
			case reqBody.Instr == "check":
				{

					t.Players[client.Name].Conn.Mut.Lock()
					t.Players[client.Name].Conn.WriteJSON(engine.ResponseData{Instr: engine.Instr.Refresh, Data: t.GetTableInfo(client.Name)})
					t.Players[client.Name].Conn.Mut.Unlock()
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
			case t.Pm.Instr != engine.Instr.Move && (reqBody.Instr == engine.Instr.PutCard || reqBody.Instr == engine.Instr.LBonus || reqBody.Instr == engine.Instr.Pass):
				fallthrough
			case !slices.Contains(t.Pm.IDs, reqBody.CardID) && reqBody.CardID != 0:
				{
					client.Conn.Mut.Lock()
					client.Conn.WriteJSON(engine.ResponseData{Instr: engine.Instr.ForbMove})
					client.Conn.Mut.Unlock()
					continue
				}
			}
			// Move validation
			err := t.MoveRouter(reqBody)
			if err != nil {
				t.Players[t.Pm.ActPlr].Conn.Mut.Lock()
				t.Players[t.Pm.ActPlr].Conn.WriteJSON(engine.ResponseData{Instr: engine.Instr.ForbMove})
				t.Players[t.Pm.ActPlr].Conn.Mut.Unlock()
				continue
			}
			ActiveGameTables.Set(t.TableID, t)
			t.RefreshTable()
			t.Players[t.Pm.ActPlr].Conn.Mut.Lock()
			t.Players[t.Pm.ActPlr].Conn.WriteJSON(engine.ResponseData{Instr: engine.Instr.Move, Data: t.Pm.IDs})
			t.Players[t.Pm.ActPlr].Conn.Mut.Unlock()
			t.Players[t.Pm.PasPlr].Conn.Mut.Lock()
			switch {
			case t.Players[t.Pm.PasPlr].PassFlag:
				{
					t.Players[t.Pm.PasPlr].Conn.WriteJSON(engine.ResponseData{Instr: engine.Instr.Pass})
				}
			case t.Pm.Instr == engine.Instr.Move:
				{
					t.Players[t.Pm.PasPlr].Conn.WriteJSON(engine.ResponseData{Instr: engine.Instr.Wait})
				}
			}
			t.Players[t.Pm.PasPlr].Conn.Mut.Unlock()

			/////////////// Move processing //////////////
		}
	}
}
