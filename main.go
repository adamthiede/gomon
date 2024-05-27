package main

import ()

type Results struct {
	CertChecks map[string]string
	PingChecks map[string]string
}

func main() {
	// temp default before reading from a config file
	config := Config{
	    CheckInterval: 1,
	    CertThreshold: 14,
	}
	parseConfig(&config)


	results := Results{
		CertChecks: map[string]string{},
		PingChecks: map[string]string{},
	}

	MonitorLoop(&config, &results)
}
