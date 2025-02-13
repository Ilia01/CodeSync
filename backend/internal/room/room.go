package room

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn    *websocket.Conn
	Room    *Room
	OutChan chan Message	
	mu sync.Mutex
}

type Room struct {
	clients   map[*Client]bool
	broadcast chan Message
	leave     chan *Client
	mu        sync.RWMutex
}

var (
	Rooms   = make(map[string]*Room)
	roomsMu sync.RWMutex
)

type Message struct {
	Type int
	Data []byte
}

func NewRoom() *Room {
	r := &Room{
		clients:   make(map[*Client]bool),
		broadcast: make(chan Message),
		leave:     make(chan *Client),
	}
	go r.Run()
	return r
}

func (r *Room) Run() {
	for {
		select {
		case message := <-r.broadcast:
			r.mu.RLock()
			for client := range r.clients {
				client.SendMessage(message)
			}
			r.mu.RUnlock()

		case client := <-r.leave:
			r.mu.Lock()
			delete(r.clients, client)
			r.mu.Unlock()
			if err := client.Conn.Close(); err != nil {
				log.Println("WebSocket close error:", err)
			}
		}
	}
}

func (r *Room) Join(client *Client) {
	r.mu.Lock()
	r.clients[client] = true
	r.mu.Unlock()
}

func (r *Room) Broadcast(msgType int, message []byte) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for client := range r.clients {
		client.SendMessage(Message{
			Type: msgType,
			Data: message,
		})
	}
}

func (r *Room) Leave(client *Client) {
	r.leave <- client
	close(client.OutChan)
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
		if err := c.Conn.Close(); err != nil {
			log.Println("WebSocket close error:", err)
		}
	}()

	for {
		msgType, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("WebSocket error: %v", err)
			}
			break
		}

		if msgType == websocket.BinaryMessage {
			log.Println("Binary message received")
		} else {
			log.Println("Received text data:", string(message))
		}	

		c.Room.Broadcast(msgType, message)
	}
}

func (c *Client) SendMessage(msg Message) {
	c.mu.Lock()
	defer c.mu.Unlock()

	select {
	case c.OutChan <- msg:
	default:
		log.Println("Client send buffer full, dropping message")
	}
}

func (c *Client) Write() {
	defer c.Conn.Close()

	for msg := range c.OutChan {
		c.mu.Lock()
		err := c.Conn.WriteMessage(msg.Type, msg.Data)
		c.mu.Unlock()

		if err != nil {
			log.Println("WebSocket write error:", err)
			return
		}
	}
}

func NewClient(conn *websocket.Conn, room *Room) *Client {
	client := &Client{
		Conn:    conn,
		Room:    room,
		OutChan: make(chan Message, 256),
	}

	go client.Write()
	return client
}