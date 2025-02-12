package main

import (
	"backend/internal/websocket"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ws/", websocket.HandleWebSocket)
	log.Println("WebSocket server started on ws://localhost:8080/ws/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
