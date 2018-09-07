package mail

import (
	"fmt"
	"net/smtp"
	"strconv"
)

type mail struct {
	fromAddr, fromPwd, host, subject, body string
	toAddr                                 []string
	port                                   int
}

// Creates a properly formatted smtp message with the to, subject and body.
func (m *mail) makeMsg() []byte {
	return []byte("To: " + m.toAddr[0] + "\r\n" +
		"Subject: " + m.subject + "\r\n" +
		"\r\n" +
		m.body + "\r\n")
}

// Sends the email
func (m *mail) SendMail() bool {
	fmt.Println("Sending email...")
	// Set up auth
	auth := smtp.PlainAuth("", m.fromAddr, m.fromPwd, m.host)
	// Create the message (formatted properly)
	msg := m.makeMsg()
	// Set the address of the mail server
	addr := m.host + ":" + strconv.Itoa(m.port)
	// Away it goes!...
	err := smtp.SendMail(addr, auth, m.fromAddr, m.toAddr, msg)
	if err != nil {
		// ...Or maybe not :(
		fmt.Println(err)
		return false
	}
	return true
}

func NewMail(fromAddr, fromPwd, host, subject, body string, toAddr []string, port int) mail {
	return mail{fromAddr, fromPwd, host, subject, body, toAddr, port}
}
