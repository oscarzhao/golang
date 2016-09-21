package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

// https://tools.ietf.org/html/rfc3501#section-8

func main() {
	addr := "imap.exmail.qq.com:143"
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Printf("error dial golang.org, error:%s\n", err)
		return
	}
	fmt.Fprintf(conn, "a001 login shuailong@tenxcloud.com xue1227xue\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Printf("error a001 login, error:%s\n", err)
		return
	}
	fmt.Printf("step 001 data:%s\n", status)

	fmt.Fprintf(conn, "a002 select inbox\r\n")
	for i := 0; i < 2; i++ {
		status, err = bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Printf("error a002 select inbox, error:%s\n", err)
			break
		}
		fmt.Printf("step 002 (%d) data:%s\n", i, status)
	}
	// fmt.Fprintf(conn, "a003 fetch 1:3 body[text]\r\n")
	fmt.Fprintf(conn, "a003 fetch 1:4 all\r\n")

	// fmt.Fprintf(conn, "a003 fetch 1:3 full\r\n")
	for i := 0; i < 7; i++ {
		status, err = bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Printf("error a003 fetch 12, error:%s\n", err)
			return
		}

		fmt.Printf("step 003 (%d) fetch 12, data:%s\n", i, status)
	}

	fmt.Printf("step 003 fetch 12(complete)\n")

	fmt.Fprintf(conn, "a004 fetch 12 body[header]\r\n")

	for i := 0; i < 2; i++ {
		status, err = bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Printf("error a004 fetch 12 body[header], error:%s\n", err)
			return
		}
		fmt.Printf("step 004 %d: %s\n", i, status)
	}

	fmt.Fprintf(conn, "a005 store 12 +flags \\seen\r\n")
	for i := 0; i < 2; i++ {
		status, err = bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Printf("error a005 store, error:%s\n", err)
			return
		}
		fmt.Printf("step 005 %d: %s\n", i, status)
	}

	fmt.Fprintf(conn, "a006 logout\r\n")
	for i := 0; i < 2; i++ {
		status, err = bufio.NewReader(conn).ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Printf("error a006 logout, error:%s\n", err)
			return
		}
		fmt.Printf("step 006 %d: %s\n", i, status)
	}

}

/*

      ENVELOPE
         A parenthesized list that describes the envelope structure of a
         message.  This is computed by the server by parsing the
         [RFC-2822] header into the component parts, defaulting various
         fields as necessary.

         The fields of the envelope structure are in the following
         order: date, subject, from, sender, reply-to, to, cc, bcc,
         in-reply-to, and message-id.  The date, subject, in-reply-to,
         and message-id fields are strings.  The from, sender, reply-to,
         to, cc, and bcc fields are parenthesized lists of address
         structures.

         An address structure is a parenthesized list that describes an
         electronic mail address.  The fields of an address structure
         are in the following order: personal name, [SMTP]
         at-domain-list (source route), mailbox name, and host name.

         [RFC-2822] group syntax is indicated by a special form of
         address structure in which the host name field is NIL.  If the
         mailbox name field is also NIL, this is an end of group marker
         (semi-colon in RFC 822 syntax).  If the mailbox name field is
         non-NIL, this is a start of group marker, and the mailbox name
         field holds the group name phrase.

         If the Date, Subject, In-Reply-To, and Message-ID header lines
         are absent in the [RFC-2822] header, the corresponding member
         of the envelope is NIL; if these header lines are present but
         empty the corresponding member of the envelope is the empty
         string.

* 12 FETCH

(FLAGS (\Seen) RFC822.SIZE 6237 INTERNALDATE "8-Aug-2016 01:15:40 +0800"
  ENVELOPE (
    "Mon, 8 Aug 2016 01:15:40 +0800" "=?GBK?B?z7XNs9Lss6O+r7jm?="    # date, subject
    (("=?GBK?B?yrHL2dTGzcW20w==?=" NIL "shuailong" "tenxcloud.com"))
    (("=?GBK?B?yrHL2dTGzcW20w==?=" NIL "shuailong" "tenxcloud.com"))
    (("=?GBK?B?yrHL2dTGzcW20w==?=" NIL "shuailong" "tenxcloud.com"))
    (("=?GBK?B?c2h1YWlsb25n?=" NIL "shuailong" "tenxcloud.com"))
    NIL
    NIL
    NIL
    ""
  )
)

*/

// https://tools.ietf.org/html/rfc3501#section-6.4.5
