package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

type Results struct {
	CertChecks map[string]string
	PingChecks map[string]string
	TcpChecks  map[string]string
	UdpChecks  map[string]string
	mux        *sync.Mutex
}

func WebServer(config *Config) {

	port := ":1314"
	if config.Port != 0 {
		port = ":" + fmt.Sprint(config.Port)
	}

	http.HandleFunc("/", getRoot)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("error serving on port: %s\n", err)
	}

	time.Sleep(time.Minute * time.Duration(config.CheckInterval))
	// infinite loop
}

func constructHtmlData() string {
	htmlHead := "<html><meta http-equiv='refresh' content='600'><meta name='color-scheme' content='dark light'><style>html { font-family : monospace }</style><title>gomon</title><body>"
	htmlFoot := "</body></html>"

	timestr := time.Now()
	htmlData := htmlHead + "<h1>Hello, world!</h1>" + "<p>" + fmt.Sprint(timestr) + "</p>" + htmlFoot
	return htmlData
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	htmlData := constructHtmlData()

	io.WriteString(w, htmlData)
}
