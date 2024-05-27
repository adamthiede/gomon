package main

import (
	"errors"
	"github.com/prometheus-community/pro-bing"
)

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
