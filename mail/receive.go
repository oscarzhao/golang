package mail

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/mail"
	"time"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
)

type MailClient struct {
	addr string
	user string
	pass string
}

// Email is a complete mail
type Email struct {
	Date    time.Time `json:"date"`
	From    string    `json:"from"`
	To      string    `json:"to"`
	Subject string    `json:"subject"`
	Content string    `json:"content"`
}

func (mc *MailClient) ListMailBox() ([]imap.MailboxInfo, error) {
	// Connect to server
	c, err := client.DialTLS(mc.addr, nil)
	if err != nil {
		return nil, err
	}
	// Login
	if err := c.Login(mc.user, mc.pass); err != nil {
		return nil, err
	}
	defer c.Logout()

	// List mailboxes
	mailboxes := make(chan *imap.MailboxInfo)
	done := make(chan error, 1)
	go func() {
		done <- c.List("", "*", mailboxes)
	}()

	var boxes []imap.MailboxInfo
	for m := range mailboxes {
		boxes = append(boxes, *m)
	}

	return boxes, <-done
}

func (mc *MailClient) ReceiveRaw(mailbox string, n uint32) ([]imap.Message, error) {
	// Connect to server
	c, err := client.DialTLS(mc.addr, nil)
	if err != nil {
		return nil, err
	}

	// Login
	if err := c.Login(mc.user, mc.pass); err != nil {
		return nil, err
	}
	defer c.Logout()

	// Select INBOX
	mbox, err := c.Select(mailbox, false)
	if err != nil {
		return nil, err
	}
	log.Printf("Flags for %s:%v\n", mailbox, mbox.Flags)

	// Get the last 4 messages
	seqset, _ := imap.NewSeqSet("")
	var start uint32
	if mbox.Messages >= n {
		start = mbox.Messages - n + 1
	} else {
		start = 1
	}
	seqset.AddRange(start, mbox.Messages)

	messages := make(chan *imap.Message)
	done := make(chan error, 1)
	go func() {
		// done <- c.Fetch(seqset, []string{imap.EnvelopeMsgAttr, imap.BodyMsgAttr, imap.SizeMsgAttr}, messages)
		done <- c.Fetch(seqset, []string{ /*imap.BodyMsgAttr, */ imap.EnvelopeMsgAttr, imap.BodyStructureMsgAttr, "BODY.PEEK[]<0.8>"}, messages)
	}()

	var msgs []imap.Message
	for msg := range messages {
		if msg.Envelope != nil {
			if bytes, err := decode(msg.Envelope.Subject); err == nil {
				msg.Envelope.Subject = string(bytes)
			}
			msg.Envelope.From = decodeArrs(msg.Envelope.From)
			msg.Envelope.To = decodeArrs(msg.Envelope.To)
			msg.Envelope.Sender = decodeArrs(msg.Envelope.Sender)
			msg.Envelope.ReplyTo = decodeArrs(msg.Envelope.ReplyTo)
			msg.Envelope.Cc = decodeArrs(msg.Envelope.Cc)
		}
		msgs = append(msgs, *msg)
	}

	return msgs, <-done
}

// ListMails lists parsed mails and errors
func (mc *MailClient) ListMails(mailbox string, n uint32) ([]Email, []error) {
	rawMessages, err := mc.ReceiveRaw(mailbox, n)
	if err != nil {
		return nil, []error{err}
	}

	var emails []Email
	var errs []error
	for _, m := range rawMessages {
		for _, literal := range m.Body {
			// fmt.Printf("sectionName: %v, content length:%d\n", sectionName.String(), literal.Len())
			r := bytes.NewReader(literal.Bytes())
			m, err := mail.ReadMessage(r)
			if err != nil {
				errs = append(errs, err)
				continue
			}

			// parse header
			header := m.Header

			email := Email{
				From:    header.Get("From"),
				To:      header.Get("To"),
				Subject: header.Get("Subject"),
			}

			email.From, err = decode(email.From)
			if err != nil {
				errs = append(errs, fmt.Errorf("decode from %s failed, error:%s\n", email.From, err))
			}

			email.To, err = decode(email.To)
			if err != nil {
				errs = append(errs, fmt.Errorf("decode email to %s, failed, error:%s", email.To, err))
			}

			email.Subject, err = decode(email.Subject)
			if err != nil {
				errs = append(errs, fmt.Errorf("decode subject %s failed, error:%s\n", email.Subject, err))
			}

			email.Date, err = time.Parse(time.RFC1123Z, header.Get("Date"))
			if err != nil {
				email.Date, err = time.Parse(time.RFC1123, header.Get("Date"))
			}

			if err != nil {
				errs = append(errs, err)
				continue
			}

			contentBytes, err := ioutil.ReadAll(m.Body)
			if err != nil {
				log.Printf("fails to read mail body, error:%s\n", err)
				errs = append(errs, err)
				continue
			}

			email.Content = string(contentBytes)
			email.Content, err = decode(email.Content)
			if err != nil {
				log.Printf("fails to decode body:%s\nerror:%s\n", email.Content, err)
			}
			// fmt.Printf("email: sub:%s, content:%s\n", email.Subject, email.Content)
			emails = append(emails, email)
		}
	}
	return emails, errs
}

func decodeArrs(addrs []*imap.Address) []*imap.Address {
	var arr []*imap.Address
	for _, from := range addrs {
		if str, err := decode(from.PersonalName); err == nil {
			from.PersonalName = str
			arr = append(arr, from)
		}
	}
	return arr
}
