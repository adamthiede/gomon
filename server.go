package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type Results struct {
	CertChecks map[string]string
	PingChecks map[string]string
	TcpChecks  map[string]string
	UdpChecks  map[string]string
}

func ServeResults(config *Config, results *Results) {
	fmt.Println("Serving results.")

	http.HandleFunc("/", getRoot)
	err := http.ListenAndServe(":3333", nil)

	if err != nil {
		fmt.Printf("error serving on port: %s\n", err)
	}

	for {
		fmt.Println(time.Now())
		time.Sleep(time.Minute * time.Duration(config.CheckInterval))
	}
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}
