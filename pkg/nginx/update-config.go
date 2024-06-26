package nginx

import (
	"os"
	"path/filepath"
	"pkg/config"
	"pkg/route"
)

// updateConfig updates the nginx configuration file with the latest routes and reloads the nginx server.
func updateConfig() error {
	routes, err := route.List()
	if err != nil {
		return err
	}

	content, err := generateConfig(routes)
	if err != nil {
		return err
	}

	configFile := filepath.Join(config.GetWorkDir(), CONFIG_FILENAME)

	if err := os.WriteFile(configFile, content, 0644); err != nil {
		return err
	}

	if err := reload(); err != nil {
		return err
	}

	return nil
}
