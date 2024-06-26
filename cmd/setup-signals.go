package main

import (
	"fmt"
	"os"
	"os/signal"
	"pkg/nginx"
	"syscall"
)

// setupSignals sets up signal handling for graceful shutdown of the application.
func setupSignals() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan // Wait for termination signal

		fmt.Println("Shutdown requested...")

		if err := nginx.Stop(); err != nil {
			fmt.Println("Failed to shutdown nginx:", err)
		}

		os.Exit(0)
	}()
}
