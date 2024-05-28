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
			if err == nil {
				results.CertChecks[checks.Name] += "+"
			} else {
				results.CertChecks[checks.Name] += "-"
			}
		}
	}
}

func monitorPing(config *Config, results *Results) {
	for _, host := range config.PingChecks {
		check, err := CheckPing(host)
		fmt.Println("ping:", host, check, err)
		if err == nil {
			results.CertChecks[host] += "+"
		} else {
			results.CertChecks[host] += "-"
		}
	}
}

func monitorTcpPort(config *Config, results *Results) {
	for _, checks := range config.TCPChecks {
		for _, port := range checks.Ports {
			check, err := CheckTcp(checks.Name, port)
			if err == nil {
				fmt.Println("TCP:", checks.Name, port, check)
				results.TcpChecks[checks.Name] += "+"
			} else {
				fmt.Println("TCP:", checks.Name, port, check, err)
				results.TcpChecks[checks.Name] += "-"
			}
		}
	}
}

func MonitorLoop(config *Config, results *Results) {
	for {
		truncateResults(results)
		go monitorPing(config, results)
		go monitorCerts(config, results)
		go monitorTcpPort(config, results)

		time.Sleep(time.Minute * time.Duration(config.CheckInterval))
	}
}

func truncateResults(results *Results) {
}
