package http_server

import (
	"net/http"
	"pkg/config"
)

func Init() error {
	http.HandleFunc("GET /_route", routeList)
	http.HandleFunc("POST /_route", routeAdd)
	http.HandleFunc("DELETE /_route/{path}", routeRemove)

	if err := http.ListenAndServe(":"+config.GetRoutesPort(), nil); err != nil {
		return err
	}

	return nil
}
