package nginx

import (
	"pkg/route"
)

// ListRoutes returns a list of routes.
func ListRoutes() ([]route.Route, error) {
	return route.List()
}
