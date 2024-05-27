package main

import (
	"fmt"
	"time"
)

func monitorCerts(config *Config, results *Results) {
	for host, port := range config.certChecks {
		check, err := CheckCertificate(host, port, config.certThreshold)
		fmt.Println("Cert:", host, port, check, err)
	}
}

func monitorPing(config *Config, results *Results) {
	for _, host := range config.pingChecks {
		check, err := CheckPing(host)
		fmt.Println("ping:", host, check, err)
	}
}

func MonitorLoop(config *Config, results *Results) {
    for {
	fmt.Println(time.Now())
	go monitorPing(config, results)
	go monitorCerts(config, results)

	time.Sleep(time.Minute * time.Duration(config.checkInterval))
    }
}
