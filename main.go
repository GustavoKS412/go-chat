package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

type templateHandler struct {
	once     sync.Once
	filename string
	template *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.template = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	t.template.Execute(w, r)
}
func main() {
	var addr = flag.String("addr", "localhost:8080", "http service address")
	flag.Parse()
	r := newRoom()

	http.Handle("/room", r)
	go r.run()

	log.Printf("Starting web server on %s", *addr)

	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("Listen and serve:", err)
	}
}
