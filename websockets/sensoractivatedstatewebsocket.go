package websockets

import (
	"net/http"
	"sync"

	"github.com/aPonce2001/wlmis-web-server/data"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgraderSensor = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clientsSensor = make(map[*websocket.Conn]bool)
var broadcastSensorState = make(chan bool)
var muSensor sync.Mutex

func handleSensorConnections(c *gin.Context) {
	w, r := c.Writer, c.Request
	conn, err := upgraderSensor.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
		return
	}

	muSensor.Lock()
	clientsSensor[conn] = true
	muSensor.Unlock()

	lastState := data.GetSensorActivatedState()

	err = conn.WriteJSON(lastState)
	if err != nil {
		muSensor.Lock()
		delete(clientsSensor, conn)
		muSensor.Unlock()
		conn.Close()
		return
	}

	for {
		var msg interface{}
		err := conn.ReadJSON(&msg)
		if err != nil {
			muSensor.Lock()
			delete(clientsSensor, conn)
			muSensor.Unlock()
			conn.Close()
			break
		}
	}
}

func handleSensorMessages() {
	for {
		state := <-broadcastSensorState
		muSensor.Lock()
		for client := range clientsSensor {
			err := client.WriteJSON(state)
			if err != nil {
				client.Close()
				delete(clientsSensor, client)
			}
		}
		muSensor.Unlock()
	}
}

func StartSensorWebSocket(router *gin.Engine) {
	router.GET("/ws/sensor-activated-state", func(c *gin.Context) {
		handleSensorConnections(c)
	})

	go handleSensorMessages()
}

func BroadcastSensorActivatedState(state bool) {
	broadcastSensorState <- state
}
