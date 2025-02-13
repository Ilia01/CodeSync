package websocket

import (
	"log"
	"net/http"
	"strings"

	"backend/internal/room"

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

	roomID := strings.TrimPrefix(r.URL.Path, "/ws/")
	if roomID == "" {
		log.Println("Invalid room ID")
		conn.Close()
		return
	}

	rm := room.GetOrCreateRoom(roomID)

	client := room.NewClient(conn, rm)
	rm.Join(client)

	go client.Read()
	go client.Write()
}
