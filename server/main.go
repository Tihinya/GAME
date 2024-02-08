package main

import (
	"log"
	"net/http"

	"bomberman-dom/engine"
	"bomberman-dom/socket"
)

func main() {
	socket.Instance = socket.NewManager()
	engine.SetBroadcaster(socket.Instance.Lobby)

	http.HandleFunc("/ws", socket.Instance.ServeWS)

	log.Println("Ctrl + Click on the link: http://localhost:8080")
	log.Println("To stop the server press `Ctrl + C`")

	http.ListenAndServe(":8080", nil)
}
