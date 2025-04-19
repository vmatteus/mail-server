package main

import (
	"net/smtp"
)

func main() {
	//auth := smtp.PlainAuth("", "", "", "localhost")

	err := smtp.SendMail("localhost:1025", nil,
		"teste@meuemail.com", []string{"vmatteus@gmail.com"},
		[]byte("Subject: Teste\r\n\r\nIsso é um teste"))
	if err != nil {
		panic(err)
	}
}
