package main

import (
	"fmt"
)

func MonitorLoop(config *Config) {
	// read config
	for host, port := range config.certChecks {
		check, err := CheckCertificate(host, port, config.certThreshold)
		fmt.Println(host, port, check, err)
	}

}
