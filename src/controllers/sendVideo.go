package controllers

import (
	"log"
	"net/http"
)

func SendVideo(w http.ResponseWriter, r *http.Request) {
	log.Println("new connection from", r.URL.Path)
	ws, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer ws.Close()
	for {
		_, w, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
			delete(Clients, ws)
			return
		}

		Video <- w
	}
}
