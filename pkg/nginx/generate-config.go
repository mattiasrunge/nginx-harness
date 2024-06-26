package nginx

import (
	"bytes"
	_ "embed"
	"path/filepath"
	"pkg/config"
	"pkg/route"
	"text/template"
)

//go:embed templates/nginx.conf
var configTemplate string

// generateConfig generates the nginx configuration file based on the provided routes.
// If any error occurs during the generation process, an error is returned.
func generateConfig(routes []route.Route) ([]byte, error) {
	tmpl, err := template.New("nginxConfig").Parse(configTemplate)
	if err != nil {
		return nil, err
	}

	params := map[string]any{
		"Routes":        routes,
		"ErrorLogFile":  filepath.Join(config.GetWorkDir(), ERROR_LOG_FILE),
		"AccessLogFile": filepath.Join(config.GetWorkDir(), ACCESS_LOG_FILE),
		"PidFile":       filepath.Join(config.GetWorkDir(), PID_FILE),
		"RoutesPort":    config.GetRoutesPort(),
		"Port":          config.GetPort(),
	}

	var b bytes.Buffer
	if err := tmpl.Execute(&b, params); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}
