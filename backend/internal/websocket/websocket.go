package websocket

import (
	"backend/internal/room"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}

	roomID := r.URL.Path[len("/ws/"):]
	rm := room.GetOrCreateRoom(roomID)

	client := room.NewClient(conn, rm)
	rm.Join(client)

	go client.Read()
	go client.Write()
}

