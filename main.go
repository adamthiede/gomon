package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
)

func main() {
	// set default config file path
	homeDir := os.Getenv("HOME")
	defaultConfigPath := homeDir + "/.config/gomon.toml"
	configPath := flag.String("config", defaultConfigPath, "path to config file")
	flag.Parse()

	// temp default before reading from a config file
	config := Config{
		CheckInterval: 1,
		CertThreshold: 14,
	}

	// parse the config
	err := parseConfig(&config, configPath)
	if err != nil {
		fmt.Printf("Error parsing config: %s\n", err)
		os.Exit(1)
	}

	results := Results{
		CertChecks: map[string]string{},
		PingChecks: map[string]string{},
		TcpChecks:  map[string]string{},
		mux:        &sync.Mutex{},
	}

	go WebServer(&config, &results)
	MonitorLoop(&config, &results)
}
