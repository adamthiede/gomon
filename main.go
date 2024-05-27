package main

import ()

type Config struct {
	certChecks    map[string][]string
	tcpChecks     map[string][]string
	pingChecks    []string
	certThreshold int
	checkInterval int
}

type Results struct {
	certChecks map[string]string
	pingChecks map[string]string
}

func main() {
	// temp default before reading from a config file
	config := Config{
		certChecks: map[string][]string{
			"adamthiede.com": {"443"},
			"github.com":     {"443"},
		},
		tcpChecks: map[string][]string{
			"192.168.7.11": {"80", "53", "8080"},
		},
		pingChecks: []string{
			"elagost.com", "postmarketos.org",
		},
		certThreshold: 14,
		checkInterval: 2,
	}

	results := Results{
		certChecks: map[string]string{},
		pingChecks: map[string]string{},
	}

	MonitorLoop(&config, &results)
}
