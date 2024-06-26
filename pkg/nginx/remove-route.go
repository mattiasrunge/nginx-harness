package nginx

import (
	"pkg/route"
)

// RemoveRoute removes a route from the nginx configuration based on the given path.
// If an error occurs during the removal or configuration update, it is returned.
func RemoveRoute(path string) error {
	if err := route.Remove(path); err != nil {
		return err
	}

	return updateConfig()
}
