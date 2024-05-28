package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
)

type Config struct {
	// with help from https://xuri.me/toml-to-go/
	// I wrote the config, parsed it in that website.
	CheckInterval int `toml:"CheckInterval"`
	CertThreshold int `toml:"CertThreshold"`
	Port          int `toml:"Port"`
	CharLimit     int `toml:"CharLimit"`
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

func parseConfig(cfg *Config, configPath *string) error {

	configData, err := os.ReadFile(*configPath)

	fmt.Println(string(configData))

	if err != nil {
		return err
	}
	_, err = toml.Decode(string(configData), cfg)
	if err != nil {
		return err
	}

	return nil
}
