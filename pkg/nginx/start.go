package nginx

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"pkg/config"
)

var proc *os.Process

// Start starts the nginx server.
// It updates the configuration, starts the nginx process as a background process,
// and prints the process ID (PID) of the started nginx process.
// It returns an error if there was a problem updating the configuration or starting the nginx process.
func Start() error {
	if err := updateConfig(); err != nil {
		return err
	}

	configFile := filepath.Join(config.GetWorkDir(), CONFIG_FILENAME)
	cmd := exec.Command("nginx", "-c", configFile)

	// Pipe stdout and stderr to the parent process
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Start nginx as a background process
	if err := cmd.Start(); err != nil {
		return err
	}

	proc = cmd.Process

	fmt.Printf("nginx started with PID %d\n", proc.Pid)

	return nil
}
