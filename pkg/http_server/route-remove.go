package http_server

import (
	"net/http"
	"pkg/nginx"
)

func routeRemove(w http.ResponseWriter, r *http.Request) {
	path := r.PathValue("path")

	if err := nginx.RemoveRoute(path); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
