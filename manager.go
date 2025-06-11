package main

import (
	"log"
	"net/http"
	"sync"
	"github.com/gorilla/websocket"
)

var (
	websocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

type Manager struct {
	clients ClientList
	sync.RWMutex
}

// a constructor/factory function to create instance of Manager
func NewManager() *Manager {
	return &Manager{
		clients: make(ClientList),
	}
}

// serveWS is a method that needs to implemented in order to establish websockets
func (m *Manager) serveWS(w http.ResponseWriter, r *http.Request) {
	log.Println("new connection")

	//upgrade regular http connection into websocket
	conn, err := websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// instantiate a new Client as soon as websocket is opened
	client := NewClient(conn, m)
	m.addClient(client)


	//Start client processss
	go client.readMessages() // reading messages concurrently on a go routine
	go client.writeMessages()
}

// takes Client as an argument and changes a client's bool to true in ClientList
func (m *Manager) addClient(client *Client) {
	m.Lock()
	defer m.Unlock()

	m.clients[client] = true
}


func (m *Manager) removeClient(client *Client) {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.clients[client]; ok {
		client.connection.Close()
		delete(m.clients, client)
	}
}

