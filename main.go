package main

import (
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

// typ reprezentujący pojedynczy szablon
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// metoda ServeHTTP obsługuje żądania HTTP
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, nil)
}

func main() {
	r := newRoom()
	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)
	// uruchomienie pokoju rozmów
	go r.run()
	// uruchomienie serwera WWW
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
