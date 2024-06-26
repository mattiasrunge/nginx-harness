package nginx

import "syscall"

// reload function to reload nginx configuration
func reload() error {
	if proc != nil {
		return proc.Signal(syscall.SIGHUP)
	}

	return nil
}
