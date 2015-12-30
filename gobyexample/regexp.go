package gobyexample

import (
	"fmt"
	"regexp"
)

func MatchIpStrict(ip string) bool { // judge if it is a ip, not include
	pattern := `^([0-9]{1,3}\.){3}[0-9]{1,3}$`
	r := regexp.MustCompile(pattern)
	return r.Match([]byte(ip))
}

func FindSubString(input string) {
	pattern := `(\S+)([ \t\r\f]+)(\S+).+`
	re := regexp.MustCompile(pattern)
	imgStr := "gcr.io/google-containers/serve_hostname                             latest                                     4cf4c64c9b03        10 months ago       4.522 MB"
	results := re.FindStringSubmatch(imgStr)
	fmt.Printf("%s:%s\n", results[1], results[3])

	srcStr := "google-containers/volume-nfs                                  0                    "
	dddd := re.FindStringSubmatch(srcStr)
	fmt.Printf("results:%#v\n", dddd)
}
