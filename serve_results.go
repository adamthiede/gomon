package main

import (
    "fmt"
    "time"
)

type Results struct {
	CertChecks map[string]string
	PingChecks map[string]string
	TcpChecks map[string]string
	UdpChecks map[string]string
}

func ServeResults(config *Config, results *Results) {
    fmt.Println("Serving results.")
    for {
	fmt.Println(time.Now())
	time.Sleep(time.Minute * time.Duration(config.CheckInterval))
    }
}
