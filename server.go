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

	port := ":1314"
	if config.Port != 0 {
		port = ":" + fmt.Sprint(config.Port)
	}

	http.HandleFunc("/", getRoot)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("error serving on port: %s\n", err)
	}

	for {
		fmt.Println(time.Now())
		time.Sleep(time.Minute * time.Duration(config.CheckInterval))
	}
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	htmlHead := "<html><meta http-equiv='refresh' content='600'><meta name='color-scheme' content='dark light'><style>html { font-family : monospace }</style><title>gomon</title><body>"
	htmlFoot := "</body></html>"

	htmlData := htmlHead + "<h1>Hello, world!</h1>" + htmlFoot

	io.WriteString(w, htmlData)
}
