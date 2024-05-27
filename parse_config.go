package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
)

type Config struct {
	CheckInterval int `toml:"CheckInterval"`
	CertThreshold int `toml:"CertThreshold"`
	CertChecks    []struct {
		Name  string   `toml:"name"`
		Ports []string `toml:"ports"`
	} `toml:"CertChecks"`
	TCPChecks []struct {
		Name  string   `toml:"name"`
		Ports []string `toml:"ports"`
	} `toml:"TcpChecks"`
	PingChecks []string `toml:"PingChecks"`
}

func parseConfig(cfg *Config) error {
	configFileName := "config.toml"

	configData, err := os.ReadFile(configFileName)

	fmt.Println(string(configData))

	if err != nil {
		fmt.Println("Cannot read config file!")
		return err
	}
	_, err = toml.Decode(string(configData), cfg)
	if err != nil {
		fmt.Println("Cannot parse config file.")
		return err
	}

	return nil
}
