package lib

import (
	"fmt"
	"net/smtp"
	"os"
	"strings"
)

var (
	hostname = "smtp.gmail.com"
	username = os.Getenv("GMAIL_USERNAME")
	password = os.Getenv("GMAIL_PASSWORD")
)

func SendGmail(to []string) {

	subject := "障害通知"
	body := "障害を検出しました"

	auth := smtp.PlainAuth("", username, password, hostname)

	msg := []byte(strings.ReplaceAll(fmt.Sprintf("To: %s\nSubject: %s\n\n%s", strings.Join(to, ","), subject, body), "\n", "\r\n"))
	if err := smtp.SendMail("smtp.gmail.com:587", auth, username, to, msg); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
