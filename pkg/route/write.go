package route

import (
	"encoding/json"
	"os"
	"path/filepath"
	"pkg/config"
)

// write writes the given routes to a file in JSON format.
func write(routes []Route) error {
	routesFile := filepath.Join(config.GetWorkDir(), ROUTES_FILENAME)

	content, err := json.Marshal(routes)
	if err != nil {
		return err
	}

	if err = os.WriteFile(routesFile, content, os.ModePerm); err != nil {
		return err
	}

	return nil
}
