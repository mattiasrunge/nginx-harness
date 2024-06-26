package route

import "slices"

// Add adds a new route to the list of routes.
// If a route with the same path already exists, it will be replaced.
// It returns an error if the update operation fails.
func Add(route Route) error {
	return update(func(routes []Route) []Route {
		routes = slices.DeleteFunc(routes, func(r Route) bool {
			return r.Path == route.Path
		})

		return append(routes, route)
	})
}
