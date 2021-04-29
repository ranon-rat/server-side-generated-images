package controllers

import (
	"log"
	"net/http"
)

func ReceiveVideo(w http.ResponseWriter, r *http.Request) {
	log.Println("new connection from", r.URL.Path)
	ws, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer ws.Close()
	Clients[ws] = true
	for {
		for client := range Clients {
			if err := client.WriteMessage(1, <-Video); err != nil {
				log.Println(err)
				delete(Clients, client)
				return
			}
		}
	}

}
