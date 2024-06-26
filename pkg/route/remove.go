package route

import "slices"

// Remove removes a route with the specified path from the routes slice.
// It returns an error if the update operation fails.
func Remove(path string) error {
	return update(func(routes []Route) []Route {
		return slices.DeleteFunc(routes, func(r Route) bool {
			return r.Path == path
		})
	})
}
