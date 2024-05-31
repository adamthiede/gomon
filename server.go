package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

type Results struct {
	CertChecks   map[string]string
	PingChecks   map[string]string
	TcpChecks    map[string]string
	UdpChecks    map[string]string
	LastCheck    string
	ErroredItems []string
	mux          *sync.Mutex
}

func WebServer(config *Config, results *Results) {

	port := ":1314"
	if config.Port != 0 {
		port = ":" + fmt.Sprint(config.Port)
	}

	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, DataParser(*results))
	}

	http.HandleFunc("/", handleFunc)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("error serving on port: %s\n", err)
	}
	fmt.Println("Server exiting!")
}

func DataParser(results Results) string {
	fmt.Println("parsing results to build html.")
	// Basic construction of page from a bunch of html and the results
	// struct. I know this isn't great, but it works fine and I like the
	// way it looks in the end.

	results.mux.Lock()
	defer results.mux.Unlock()

	htmlHead := `<html>
	<meta http-equiv='refresh' content='600'>
	<meta name='color-scheme' content='dark light'>
	<style>html { font-family : monospace }</style>
	<title>gomon</title>
	<body>`
	htmlHead += "<p> Time:" + fmt.Sprint(time.Now()) + "</p>"
	htmlHead += "<p> Last check:" + results.LastCheck + "</p>"
	htmlFoot := "</body></html>"

	htmlCerts := "<h2>Cert checks:</h2><p>"
	htmlTcp := "<h2>Tcp checks:</h2><p>"
	htmlPing := "<h2>Ping checks:</h2><p>"

	for k, v := range results.CertChecks {
		htmlCerts += "<p>" + k + ": " + v + "</p>"
	}
	htmlCerts += "</p>"

	for k, v := range results.TcpChecks {
		htmlTcp += fmt.Sprintf("<p>%s: %s</p>\n", k, v)
	}
	htmlTcp += "</p>"

	for k, v := range results.PingChecks {
		htmlPing += "<p>" + k + ": " + v + "</p>"
	}
	htmlPing += "</p>"

	return htmlHead + htmlCerts + htmlTcp + htmlPing + htmlFoot
}
