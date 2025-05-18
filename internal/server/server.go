package server

import (
	"log"
	"net/http"

	"github.com/apbenedetti/maple-pie/internal/api"
)

func Start() {

	fs := http.FileServer(http.Dir("./ui/static"))
	http.Handle("/", fs)

	http.HandleFunc("/api/opencanada", api.OpenCanadaHandler)
	http.Handle("/config/", http.StripPrefix("/config/", http.FileServer(http.Dir("./ui/config"))))

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
