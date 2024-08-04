package websockets

import (
	"net/http"
	"sync"
	"time"

	"github.com/aPonce2001/wlmis-web-server/data"
	"github.com/aPonce2001/wlmis-web-server/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type getWaterLevelRecordJSON struct {
	HeightCm float64   `json:"heightCm"`
	Percent  float64   `json:"percent"`
	VolumeMl float64   `json:"volumeMl"`
	DateTime time.Time `json:"dateTime"`
}

var upgraderWaterLevel = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clientsWaterLevel = make(map[*websocket.Conn]bool)
var broadcastWaterLevel = make(chan getWaterLevelRecordJSON)
var muWaterLevel sync.Mutex

func handleWaterLevelConnections(context *gin.Context) {
	writer, request := context.Writer, context.Request
	connection, err := upgraderWaterLevel.Upgrade(writer, request, nil)
	if err != nil {
		http.Error(writer, "Could not open websocket connection", http.StatusBadRequest)
		return
	}

	muWaterLevel.Lock()
	clientsWaterLevel[connection] = true
	muWaterLevel.Unlock()

	lastRecordsModel := data.GetLastNWaterLevelRecords(20)

	var lastRecords []getWaterLevelRecordJSON
	for _, record := range lastRecordsModel {
		lastRecords = append(lastRecords, getWaterLevelRecordJSON{
			HeightCm: record.HeightCm,
			Percent:  record.Percent,
			VolumeMl: record.VolumeMl,
			DateTime: record.DateTime,
		})
	}

	err = connection.WriteJSON(lastRecords)
	if err != nil {
		muWaterLevel.Lock()
		delete(clientsWaterLevel, connection)
		muWaterLevel.Unlock()
		connection.Close()
		return
	}

	for {
		var message interface{}
		err := connection.ReadJSON(&message)
		if err != nil {
			muWaterLevel.Lock()
			delete(clientsWaterLevel, connection)
			muWaterLevel.Unlock()
			connection.Close()
			break
		}
	}
}

func handleWaterLevelMessages() {
	for {
		msg := <-broadcastWaterLevel
		muWaterLevel.Lock()
		for client := range clientsWaterLevel {
			err := client.WriteJSON(msg)
			if err != nil {
				client.Close()
				delete(clientsWaterLevel, client)
			}
		}
		muWaterLevel.Unlock()
	}
}

func StartWaterLevelWebSocket(router *gin.Engine) {
	router.GET("/ws/water-level", func(context *gin.Context) {
		handleWaterLevelConnections(context)
	})

	go handleWaterLevelMessages()
}

func BroadcastWaterLevel(record models.WaterLevelRecord) {
	jsonRecord := getWaterLevelRecordJSON{
		HeightCm: record.HeightCm,
		Percent:  record.Percent,
		VolumeMl: record.VolumeMl,
		DateTime: record.DateTime,
	}
	broadcastWaterLevel <- jsonRecord
}
