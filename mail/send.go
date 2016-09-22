/*
  used to send mail
*/
package mail

import (
	"fmt"
	"net/smtp"
	"strings"
)

const (
	ImapHost       = "imap.exmail.qq.com" // receive mail
	ImapPortSecure = 993                  //receive mail
	Smtphost       = "smtp.exmail.qq.com" // send mail
	Smtpport       = 25                   // send mail
	Identity       = "Oscar"
	User           = "shuailong@tenxcloud.com"
	Pass           = "xxx"
)

// SendMail enables a user to send mail to any mails
func SendMail(to []string, subject, body string) error {
	// Set up authentication information.
	auth := smtp.PlainAuth(Identity, User, Pass, Smtphost)

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	// msg := fmt.Sprintf("To: %s\r\nFrom: %s<%s>\r\nSubject: %s\r\nContent-Type: text/html; charset=UTF-8\r\n\r\n%s\r\n", strings.Join(to, ";"), identity, user, subject, body)
	msg := fmt.Sprintf("To: %s\r\n", strings.Join(to, ";")) +
		fmt.Sprintf("From: %s<%s>\r\n", Identity, User) +
		fmt.Sprintf("Subject: %s\r\n", subject) +
		"Content-Type: text/html; charset=UTF-8\r\n\r\n" +
		fmt.Sprintf("%s\r\n", body)
	fmt.Printf("msg: %s\n", msg)
	return smtp.SendMail(fmt.Sprintf("%s:%d", Smtphost, Smtpport), auth, User, to, []byte(msg))
}
