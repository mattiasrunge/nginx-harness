package nginx

import (
	"fmt"
	"syscall"
)

// Stop stops the nginx process gracefully.
// It sends a SIGTERM signal to the process and waits for it to exit.
// If the process is not running, it does nothing.
// Returns an error if there was a problem sending the signal or waiting for the process to exit.
func Stop() error {
	if proc != nil {
		fmt.Println("Shutting down nginx...")

		if err := proc.Signal(syscall.SIGTERM); err != nil {
			return err
		}

		state, err := proc.Wait()

		if err != nil {
			return err
		}

		fmt.Println("nginx stopped with state:", state)
	}

	return nil
}
