package main

import (
	"flag"
	"fmt"
	"os"
	"pkg/config"
	"pkg/http_server"
	"pkg/nginx"
)

func main() {
	fmt.Println("Starting nginx-harness...")

	setupSignals()

	workDirPtr := flag.String("dir", "", "Path to the work directory")
	portPtr := flag.String("port", "9898", "Port to start HTTP server on")

	flag.Parse()

	if *workDirPtr == "" {
		fmt.Println("A work directory must be specified.")
		flag.PrintDefaults()
		os.Exit(1)
	}

	config.SetWorkDir(*workDirPtr)
	config.SetPort(*portPtr)

	fmt.Printf("Working directory is %s\n", config.GetWorkDir())
	fmt.Printf("HTTP port is %s\n", config.GetPort())

	if err := nginx.Start(); err != nil {
		fmt.Println("Failed to start nginx:", err)
		os.Exit(1)
	}

	if err := http_server.Init(); err != nil {
		fmt.Println("Failed to start routes interface:", err)
		os.Exit(1)
	}

}
