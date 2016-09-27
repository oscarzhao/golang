package mail

import (
	// "encoding/json"
	"fmt"
	"testing"
)

func TestListBoxes(t *testing.T) {
	mc := MailClient{
		addr: fmt.Sprintf("%s:%d", ImapHost, ImapPortSecure),
		user: User,
		pass: Pass,
	}
	boxes, err := mc.ListMailBox()
	if err != nil {
		t.Fatalf("list mail box fails, error:%s\n", err)
	}
	for _, b := range boxes {
		t.Logf("box name: %s\n", b.Name)
	}
	// t.Logf("boxes: %#v\n", boxes)
}

// func TestReceiveRaw(t *testing.T) {
// 	num := uint32(1)
// 	mc := MailClient{
// 		addr: fmt.Sprintf("%s:%d", ImapHost, ImapPortSecure),
// 		user: User,
// 		pass: Pass,
// 	}
// 	box := "INBOX"
// 	messages, err := mc.ReceiveRaw(box, num)
// 	if err != nil {
// 		t.Fatalf("list mail box fails, error:%s\n", err)
// 	}
// 	for _, m := range messages {
// 		bytes, err := json.MarshalIndent(m.BodyStructure, "", "  ")
// 		if err != nil {
// 			t.Errorf("err marshal message: %s\n", err)
// 		} else {
// 			t.Logf("body structure: \n%s\n", bytes)
// 		}

// 		t.Logf("\nbody:\n")
// 		for sectionName, literal := range m.Body {
// 			t.Logf("sectionName: %v, value len:%d, contents: %s\n", sectionName.String(), literal.Len(), literal.Bytes())
// 		}

// 		bytes, err = json.MarshalIndent(m.Envelope, "", "  ")
// 		if err != nil {
// 			t.Errorf("err marshal message: %s\n", err)
// 		} else {
// 			t.Logf("Envelope : \n%s\n", bytes)
// 		}

// 		t.Logf("seq num:%#v\n", m.SeqNum)
// 		t.Logf("\nseq number: %d\nsubject: %s\n\n", m.SeqNum, m.Envelope.Subject)
// 	}
// 	// t.Logf("messages: %#v\n", messages)
// }

func TestListMails(t *testing.T) {
	num := uint32(9)
	mc := MailClient{
		addr: fmt.Sprintf("%s:%d", ImapHost, ImapPortSecure),
		user: User,
		pass: Pass,
	}
	box := "INBOX"
	messages, errs := mc.ListMails(box, num)
	if errs != nil {
		t.Fatalf("list mail box fails, error:%v\n", errs)
	}
	for _, m := range messages {
		t.Logf("\nSubject: %s, From: %s, Date:%s\nContent:\n%s\n", m.Subject, m.From, m.Date, m.Content)
	}
}
