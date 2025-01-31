package server

import (
	"net/http"

	"github.com/apbenedetti/maple-pie/internal/api"
)

func Start() {

	http.HandleFunc("/", api.Handler)

	http.ListenAndServe(":8080", nil)
}
