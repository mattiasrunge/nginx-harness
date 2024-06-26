package http_server

import (
	"encoding/json"
	"net/http"
	"pkg/nginx"
	"pkg/route"
)

func routeAdd(w http.ResponseWriter, r *http.Request) {
	newRoute := route.Route{}

	if err := json.NewDecoder(r.Body).Decode(&newRoute); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := nginx.AddRoute(newRoute); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
