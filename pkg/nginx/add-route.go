package nginx

import (
	"pkg/route"
)

// AddRoute adds a new route to the nginx configuration.
// It takes a route.Route object as a parameter and returns an error if any.
func AddRoute(r route.Route) error {
	if err := route.Add(r); err != nil {
		return err
	}

	return updateConfig()
}
