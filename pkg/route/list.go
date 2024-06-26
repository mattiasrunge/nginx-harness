package route

import (
	"encoding/json"
	"os"
	"path/filepath"
	"pkg/config"
)

// List reads routes from a file and returns them as a slice of Route structs.
func List() ([]Route, error) {
	routesFile := filepath.Join(config.GetWorkDir(), ROUTES_FILENAME)

	_, err := os.Stat(routesFile)
	if os.IsNotExist(err) {
		return []Route{}, nil
	}

	data, err := os.ReadFile(routesFile)
	if err != nil {
		return nil, err
	}

	var routes []Route
	if err = json.Unmarshal(data, &routes); err != nil {
		return nil, err
	}

	return routes, nil
}
