package controllers

import "github.com/gorilla/websocket"

var (
	Clients  = make(map[*websocket.Conn]bool)
	Video    = make(chan []byte)
	Upgrader = websocket.Upgrader{}
)
