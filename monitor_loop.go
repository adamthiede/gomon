package main

import (
	"fmt"
	"time"
)

func monitorCerts(config *Config, results *Results) {
	results.mux.Lock()
	for _, checks := range config.CertChecks {
		for _, port := range checks.Ports {
			check, err := CheckCertificate(checks.Name, port, config.CertThreshold)
			fmt.Println("Cert:", checks.Name, port, check, err)
			checkName := fmt.Sprintf("%s %v", checks.Name, port)
			if err == nil {
				results.CertChecks[checkName] += "+"
			} else {
				results.CertChecks[checkName] += "-"
			}
		}
	}
	results.mux.Unlock()
}

func monitorPing(config *Config, results *Results) {
	results.mux.Lock()
	for _, host := range config.PingChecks {
		check, err := CheckPing(host)
		fmt.Println("ping:", host, check, err)
		if err == nil {
			results.PingChecks[host] += "+"
		} else {
			results.PingChecks[host] += "-"
		}
	}
	results.mux.Unlock()
}

func monitorTcpPort(config *Config, results *Results) {
	results.mux.Lock()
	for _, checks := range config.TCPChecks {
		for _, port := range checks.Ports {
			check, err := CheckTcp(checks.Name, port)
			checkName := checks.Name + ":" + port
			if err == nil {
				fmt.Println("TCP:", checkName, check)
				results.TcpChecks[checkName] += "+"
			} else {
				fmt.Println("TCP:", checkName, check, err)
				results.TcpChecks[checkName] += "-"
			}
		}
	}
	results.mux.Unlock()
}

func MonitorLoop(config *Config, results *Results) {
	for {
		go monitorPing(config, results)
		go monitorCerts(config, results)
		go monitorTcpPort(config, results)
		go DataParser(*results)
		time.Sleep(time.Minute * time.Duration(config.CheckInterval))
	}
}
