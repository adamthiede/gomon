package main

import (
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/prometheus-community/pro-bing"
	"net"
	"time"
)

const timeout = time.Second * 3

func CheckCertificate(host string, port string, threshold int) (bool, error) {
	// these values should be configurable
	defaultPort := "443"
	certExpiryThreshold := (time.Hour * 24) * time.Duration(threshold)

	if port == "" {
		port = defaultPort
	}

	hostToCheck := host + ":" + port

	d := net.Dialer{Timeout: timeout}
	conn, err := tls.DialWithDialer(&d, "tcp", hostToCheck, nil)
	if err != nil {
		certError := fmt.Sprintf("TLS not supported: %s", hostToCheck)
		return false, errors.New(certError)
	}

	err = conn.VerifyHostname(host)
	if err != nil {
		verityError := fmt.Sprintf("Hostname doesn't match: %s", hostToCheck)
		return false, errors.New(verityError)
	}

	expiry := conn.ConnectionState().PeerCertificates[0].NotAfter

	if time.Until(expiry) < certExpiryThreshold {
		timeError := fmt.Sprintf("Cert is about to expire for %s (%s)", hostToCheck, time.Until(expiry))
		return false, errors.New(timeError)
	}

	conn.Close()

	return true, nil
}

func CheckPing(host string) (bool, error) {
	pinger, err := probing.NewPinger(host)
	if err != nil {
		return false, errors.New("could not create pinger")
	}
	pinger.Count = 2
	err = pinger.Run() // Blocks until finished.
	if err != nil {
		return false, errors.New("ping failed.")
	}
	return true, nil
}

func CheckTcp(host string, port string) (bool, error) {
	// these values should be configurable
	defaultPort := "22"

	if port == "" {
		port = defaultPort
	}

	hostToCheck := host + ":" + port

	d := net.Dialer{Timeout: timeout}
	conn, err := d.Dial("tcp", hostToCheck)
	if err != nil {
		dialError := fmt.Sprintf("Could not connect to port: %s:%s", host, port)
		return false, errors.New(dialError)
	}

	conn.Close()

	return true, nil
}
