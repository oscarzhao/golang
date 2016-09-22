package mail

import (
	"fmt"
	"testing"
)

func TestListBoxes(t *testing.T) {
	mc := MailClient{
		addr: fmt.Sprintf("%s:%d", ImapHost, ImapPortSecure),
		user: "xxx",
		pass: "xxxx",
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

func TestReceive(t *testing.T) {
	num := uint32(1)
	mc := MailClient{
		addr: fmt.Sprintf("%s:%d", ImapHost, ImapPortSecure),
		user: "shuailong@tenxcloud.com",
		pass: "xue1227xue",
	}
	messages, err := mc.Receive(num)
	if err != nil {
		t.Fatalf("list mail box fails, error:%s\n", err)
	}
	for _, m := range messages {
		t.Logf("seq num:%#v\n", m.SeqNum)

		decoded, err := decode(m.Envelope.Subject)
		if err != nil {
			t.Fatalf("decode subject: %s failed, error:%s\n", m.Envelope.Subject, err)
		}
		t.Logf("\nseq number: %d\nsubject: %s\ndecoded: %s\n\n", m.SeqNum, m.Envelope.Subject, decoded)
	}
	// t.Logf("messages: %#v\n", messages)
}
