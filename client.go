package main

import (
	"github.com/gorilla/websocket"
)

// Typ client reprezentuje pojedynczego użytkownika prowadzącego konwersację
// z użyciem komunikatora
type client struct {
	// socket to gniazdo internetowe do obsługi danego klienta
	socket *websocket.Conn

	// send to kanał, którym przesyłane są komunikaty
	send chan []byte

	// room to pokój rozmów uzywany przez klienta
	room *room
}

func (c *client) read() {
    defer c.socket.Close()
    for {
        _, msg, err := c.socket.ReadMessage()
        if err != nil {
            return
        }
        c.room.forward <- msg
    }
}

func (c *client) write() {
    defer c.socket.Close()
    for msg := range c.send {
        err := c.socket.WriteMessage(websocket.TextMessage, msg)
        if err != nil {
            return
        }
    }
}
