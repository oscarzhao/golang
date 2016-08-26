package main

import (
	"fmt"
	"net/http"
  "regexp"
  "strings"
)

func main() {
	invalidUrls := []string{
		"http://www.does-not-exist.oscarzhao",
		"q://q.com",
		"https://www.does-not-exist.com",
	}

	pattern := `Get (\S+)://(\S+):(.+)`
	re := regexp.MustCompile(pattern)

	for _, url := range invalidUrls {
		_, err := http.Get(url)
		errMsg := err.Error()
		results := re.FindStringSubmatch(errMsg)
    index := len(results)-1
		fmt.Printf("%#v\n", strings.Trim(results[index], " \n\t"))
	}

}
