package main

import (
	"fmt"
	"time"
)

func monitorCerts(config *Config, results *Results) {
	for _, checks := range config.CertChecks {
		for _, port := range checks.Ports {
			check, err := CheckCertificate(checks.Name, port, config.CertThreshold)
			fmt.Println("Cert:", checks.Name, port, check, err)
		}
	}
}

func monitorPing(config *Config, results *Results) {
	for _, host := range config.PingChecks {
		check, err := CheckPing(host)
		fmt.Println("ping:", host, check, err)
	}
}

func monitorTcpPort(config *Config, results *Results) {
	for _, checks := range config.TCPChecks {
		for _, port := range checks.Ports {
			check, err := CheckTcp(checks.Name, port)
			if err == nil {
				fmt.Println("TCP:", checks.Name, port, check)
			} else {
				fmt.Println("TCP:", checks.Name, port, check, err)
			}
		}
	}
}

func MonitorLoop(config *Config, results *Results) {
	for {
		fmt.Println(time.Now())
		go monitorPing(config, results)
		go monitorCerts(config, results)
		go monitorTcpPort(config, results)

		time.Sleep(time.Minute * time.Duration(config.CheckInterval))
	}
}
