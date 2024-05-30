package main

import (
	"fmt"
	"time"
)

func monitorCerts(config *Config, results *Results) {
	results.mux.Lock()
	defer results.mux.Unlock()
	for _, checks := range config.CertChecks {
		for _, port := range checks.Ports {
			check, err := CheckCertificate(checks.Name, port, config.CertThreshold)
			fmt.Println("Cert:", checks.Name, port, check, err)
			checkName := fmt.Sprintf("%s %v", checks.Name, port)
			if err == nil {
				results.CertChecks[checkName] += "+"
			} else {
				results.CertChecks[checkName] += "-"
				results.ErroredItems = append(results.ErroredItems, checkName)
			}
		}
	}
	results.LastCheck = fmt.Sprintf("%s", time.Now())
}

func monitorPing(config *Config, results *Results) {
	results.mux.Lock()
	defer results.mux.Unlock()
	for _, host := range config.PingChecks {
		check, err := CheckPing(host)
		fmt.Println("ping:", host, check, err)
		if err == nil {
			results.PingChecks[host] += "+"
		} else {
			results.PingChecks[host] += "-"
			results.ErroredItems = append(results.ErroredItems, host)
		}
	}
	results.LastCheck = fmt.Sprintf("%s", time.Now())
}

func monitorTcpPort(config *Config, results *Results) {
	results.mux.Lock()
	defer results.mux.Unlock()
	for _, checks := range config.TCPChecks {
		for _, port := range checks.Ports {
			check, err := CheckTcp(checks.Name, port)
			checkName := checks.Name + ":" + port
			if err == nil {
				fmt.Println("TCP:", checkName, check)
				results.TcpChecks[checkName] += "+"
			} else {
				fmt.Println("TCP:", checkName, check, err)
				results.TcpChecks[checkName] += "-"
			results.ErroredItems = append(results.ErroredItems, checkName)
			}
		}
	}
	results.LastCheck = fmt.Sprintf("%s", time.Now())
}

func CheckIfSendEmail(config *Config, results *Results) {
    if config.Email.To == "" {
	fmt.Println("No email setup.")
	return
    }
    fmt.Printf("Errors: %s\n",results.ErroredItems)
    if len(results.ErroredItems)>0 {
	fmt.Printf("Have %v errors, sending email\n",len(results.ErroredItems))
	alerts:="\n"
	for _, errItem := range results.ErroredItems {
	    alerts+=errItem+"\n"
	}
	err := SendEmail([]byte(alerts), config.Email.To, config.Email.From, config.Email.Server, config.Email.Port, config.Email.Password)
	if err!= nil {
	    fmt.Println(err)
	}
    }
}

func MonitorLoop(config *Config, results *Results) {
	for {
		go monitorPing(config, results)
		go monitorCerts(config, results)
		go monitorTcpPort(config, results)
		truncateGraph(config, *results)
		go DataParser(*results)
		time.Sleep(time.Minute * time.Duration(config.CheckInterval))
		CheckIfSendEmail(config, results)
	}
}
func truncateGraph(config *Config, results Results) {
	results.mux.Lock()
	defer results.mux.Unlock()
	limit := 25
	if config.CharLimit > 0 {
		limit = config.CharLimit
	}
	for k, v := range results.CertChecks {
		if len(v) > limit {
			fmt.Println("truncating c")
			results.CertChecks[k] = v[len(v)-limit : len(v)-1]
		}
	}
	for k, v := range results.TcpChecks {
		if len(v) > limit {
			fmt.Println("truncating t")
			results.TcpChecks[k] = v[len(v)-limit : len(v)-1]
		}
	}
	for k, v := range results.PingChecks {
		if len(v) > limit {
			fmt.Println("truncating p")
			results.PingChecks[k] = v[len(v)-limit : len(v)-1]
		}
	}
}
