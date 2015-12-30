package gobyexample

import (
	"fmt"
	"regexp"
  "errors"
)

func MatchIpStrict(ip string) bool { // judge if it is a ip, not include
	pattern := `^([0-9]{1,3}\.){3}[0-9]{1,3}$`
	r := regexp.MustCompile(pattern)
	return r.Match([]byte(ip))
}

/** 
 * input  docker images result rows: one row
 * e.g. "gcr.io/google-containers/serve_hostname                             latest                                     4cf4c64c9b03        10 months ago       4.522 MB"
 * output  repo:tag
 * e.g. "gcr.io/google-containers/serve_hostname:latest"
 */
func ParseDockerImage(input string) (string, error) {
	pattern := `(\S+)([ \t\r\f]+)(\S+).+`
	re := regexp.MustCompile(pattern)
	results := re.FindStringSubmatch(input)
  fmt.Println(results)
  if len(results) < 4 {
    return "", errors.New("input string is not a valid \"docker images\" output")
  }
  return fmt.Sprintf("%s:%s", results[1], results[3]), nil
}
