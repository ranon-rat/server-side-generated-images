package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ranon-rat/video-transmission/src/controllers"
)

func SetupRoutes() error {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "view/home.html")
	})
	r.HandleFunc("/admin/sendVideo", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "view/sendVideo.html")
	})
	r.HandleFunc(`/public/{file:[\w\W\/]+}`, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	// websocket stuff
	r.HandleFunc("/ws/receiveVideo/", controllers.ReceiveVideo)
	r.HandleFunc("/ws/sendVideo/", controllers.SendVideo)
	log.Println("server on http://localhost:8080")
	return http.ListenAndServe(":8080", r)
}
