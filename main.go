package main

import ()

type Config struct {
	certChecks    map[string]string
	pingChecks    []string
	certThreshold int
}

func main() {
	config := Config{
		certChecks: map[string]string{
			"adamthiede.com": "443",
			"github.com":     "443",
		},
		pingChecks: []string{
			"elagost.com", "postmarketos.org",
		},
		certThreshold: 14,
	}
	MonitorLoop(&config)
}
