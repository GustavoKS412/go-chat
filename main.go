package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	var addr = flag.String("addr", ":8080", "http address")
	flag.Parse()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/", &templateHandler{filename: "index.html"})
	http.Handle("/chat", &templateHandler{filename: "chat.html"})
	http.HandleFunc("/room", func(w http.ResponseWriter, r *http.Request) {
		roomName := r.URL.Query().Get("room")
		if roomName == "" {
			http.Error(w, "room name is required", http.StatusBadRequest)
			return
		}
		realRoom := getRoom(roomName)
		realRoom.ServeHTTP(w, r)
	})

	log.Printf("Starting web server on %s", *addr)

	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
