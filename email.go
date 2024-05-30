package main

import (
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
)

func SendEmail(alertItems []byte, toAddr string, fromAddr string, server string, port string, password string) error {
	fullServer := server + ":" + port
	e := email.NewEmail()
	e.From = fromAddr
	e.To = []string{toAddr}
	e.Subject = "[gomon] server alert"
	e.Text = []byte(alertItems)
	err := e.Send(fullServer, smtp.PlainAuth("", fromAddr, password, server))
	if err != nil {
		return err
	}
	fmt.Println("Email sent!")

	return nil
}
