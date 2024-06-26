package route

type Route struct {
	Path   string `json:"path"`
	Target string `json:"target"`
}

type Updater func(routes []Route) []Route
