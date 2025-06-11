package main

import (
	"github.com/gorilla/websocket"
	"log"
)

type ClientList map[*Client]bool // goes as an argument in manager struct


// goes as an argument ins ClientList
type Client struct {
	connection *websocket.Conn
	manager *Manager
}

// a constructor/factory function to create instance of Client
func NewClient(conn *websocket.Conn, manager *Manager) *Client {
	return &Client{
		connection: conn,
		manager: manager,
	}
}


func (c *Client) readMessages(){
	defer func() {
		//cleanup connection
		c.manager.removeClient(c)
	}()

	for {
		messageType, payload, err := c.connection.ReadMessage()

		if err != nil{
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure){
				log.Printf("error reading message: %v",err)
			}
			break
		}

		log.Println(messageType)
		log.Println(string(payload))
	}
}
