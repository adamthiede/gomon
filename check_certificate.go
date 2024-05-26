package main

import (
	"crypto/tls"
	"errors"
	"fmt"
	"time"
)

func CheckCertificate(host string, port string, threshold int) (bool, error) {
	// these values should be configurable
	defaultPort := "443"
	certExpiryThreshold := (time.Hour * 24) * time.Duration(threshold)

	if port == "" {
		port = defaultPort
	}

	hostToCheck := host + ":" + port

	conn, err := tls.Dial("tcp", hostToCheck, nil)
	if err != nil {
		certError := fmt.Sprintf("TLS not supported: %s:%s", host, defaultPort)
		return false, errors.New(certError)
	}

	err = conn.VerifyHostname(host)
	if err != nil {
		verityError := fmt.Sprintf("Hostname doesn't match: %s:%s", host, defaultPort)
		return false, errors.New(verityError)
	}

	expiry := conn.ConnectionState().PeerCertificates[0].NotAfter

	if time.Until(expiry) < certExpiryThreshold {
		timeError := fmt.Sprintf("Cert is about to expire for %s:%s (%s)", host, defaultPort, time.Until(expiry))
		return false, errors.New(timeError)
	}

	conn.Close()

	return true, nil
}
