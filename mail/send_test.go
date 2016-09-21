package mail

import (
	"testing"
)

func TestSendMail(t *testing.T) {
	to := []string{"zhaoshuailong09@163.com"}
	subject := "subject for test"
	body := `
<html>
  <head>
    <title>this is a test title<title>
  </head>
  <body>
    this is a test for sending mail
  </body>
</html>`
	if err := SendMail(to, subject, body); err != nil {
		t.Fatalf("send mail to %v fails, error:%s\n", to, err)
	}
}
