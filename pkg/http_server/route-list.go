package http_server

import (
	"encoding/json"
	"net/http"
	"pkg/nginx"
)

func routeList(w http.ResponseWriter, r *http.Request) {
	routes, err := nginx.ListRoutes()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(routes); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
