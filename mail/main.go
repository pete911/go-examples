package main

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"net/smtp"
	"os"
	"strings"
)

var (
	smtpHost = "smtp.gmail.com"
	smtpPort = 587
	smtpAddr = fmt.Sprintf("%s:%d", smtpHost, smtpPort)
)

func main() {

	var (
		from    = prompt("from email: ")
		pass    = promptPass("password: ")
		to      = []string{prompt("to email: ")}
		message = []byte(prompt("message: "))
	)

	auth := smtp.PlainAuth("", from, string(pass), smtpHost)
	if err := smtp.SendMail(smtpAddr, auth, from, to, message); err != nil {
		log.Fatalf("send mail: %v", err)
	}
}

func prompt(msg string) string {
	fmt.Print(msg)
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Fatalf("read input: %v", err)
	}
	return strings.TrimSuffix(input, "\n")
}

func promptPass(msg string) []byte {
	fmt.Print(msg)
	pass, err := terminal.ReadPassword(0)
	fmt.Println()
	if err != nil {
		log.Fatalf("read password: %v", err)
	}
	return pass
}
