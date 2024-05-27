package main

import (
	"errors"
	"fmt"
	"net"
)

func CheckTcp(host string, port string) (bool, error) {
	// these values should be configurable
	defaultPort := "22"

	if port == "" {
		port = defaultPort
	}

	hostToCheck := host + ":" + port

	conn, err := net.Dial("tcp", hostToCheck)
	if err != nil {
		dialError := fmt.Sprintf("Could not connect to port: %s:%s", host, port)
		return false, errors.New(dialError)
	}

	conn.Close()

	return true, nil
}
