package gobyexample

import (
	"bytes"
	"fmt"
	"regexp"
	"testing"
)

func TestMatchIpStrict(t *testing.T) {
	// match ip test
	should_match := []string{"127.0.0.1", "0.0.0.0", "255.255.255.255", "192.168.1.2"}
	for _, ip := range should_match {
		if !MatchIpStrict(ip) {
			t.Errorf("failure: ip %s should be A IP addr\n", ip)
		} else {
			t.Logf("ip %s is a ip addr\n", ip)
		}
	}

	should_not_match := []string{"127.0.0.1a", "a127.0.0.1", "a127.0.0.1a"}
	for _, ip := range should_not_match {
		if MatchIpStrict(ip) {
			t.Errorf("failure: ip %s should not be An IP addr\n", ip)
		} else {
			t.Logf("ip %s is not ip addr\n", ip)
		}
	}
}

func TestOthers(t *testing.T) {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch")

	fmt.Println(r.MatchString("peach"))
	fmt.Println(r.FindString("peach punch"))
	fmt.Println(r.FindStringIndex("peach punch"))
	fmt.Println(r.FindStringSubmatch("peach punch"))
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	fmt.Println(r.FindAllString("peach punch pinch", -1))
	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))
	fmt.Println(r.FindAllString("peach punch pinch", 2))
	fmt.Println(r.Match([]byte("peach")))

	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Printf("\n\n----%s---\n", string(out))
}
