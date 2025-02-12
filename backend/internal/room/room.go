package room

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn    *websocket.Conn
	Room    *Room
	OutChan chan []byte
}

type Room struct {
	clients   map[*Client]bool
	broadcast chan []byte
	leave     chan *Client
	mu        sync.RWMutex
}

var (
	Rooms   = make(map[string]*Room)
	roomsMu sync.Mutex
)

func NewRoom() *Room {
	r := &Room{
		clients:   make(map[*Client]bool),
		broadcast: make(chan []byte),
		leave:     make(chan *Client),
	}
	go r.Run()
	return r
}

func (r *Room) Run() {
	for {
		select {
		case message := <-r.broadcast:
			r.Broadcast(message)
		case client := <-r.leave:
			r.Leave(client)
		}
	}
}

func (r *Room) Join(client *Client) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.clients[client] = true
}

func (r *Room) Broadcast(message []byte) {
	r.broadcast <- message
}

func (r *Room) Leave(client *Client) {
	r.leave <- client
}

func GetOrCreateRoom(roomID string) *Room {
	roomsMu.Lock()
	defer roomsMu.Unlock()

	if r, exists := Rooms[roomID]; exists {
		return r
	}

	r := NewRoom()
	Rooms[roomID] = r
	return r
}

func (c *Client) Read() {
	defer func() {
		c.Room.Leave(c)
		c.Conn.Close()
	}()

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("WebSocket error: %v", err)
			}
			break
		}
		c.Room.Broadcast(message)
	}
}

func (c *Client) SendMessage(message []byte) {
	select {
	case c.OutChan <- message:
	default:
		log.Println("Client send buffer full, dropping message")
	}
}

func (c *Client) Write() {
	defer c.Conn.Close()

	for message := range c.OutChan {
		if err := c.Conn.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Println("WebSocket write error:", err)
			return
		}
	}
}

func  NewClient(conn *websocket.Conn, room *Room) *Client {
	return &Client{
		Conn:    conn,
		Room:    room,
		OutChan: make(chan []byte, 256),
	}
}

