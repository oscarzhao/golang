package mail

import (
	"mime"
	"strings"
)

// decode is used to decode utf-8 encoding with quoted printable string to plain text
// warning: if the plain text is too long, it will be cut and decoded and assembled together(look at the test file)
// warning: mail content does NOT have encoding prefix and ?= suffix
// http://superuser.com/questions/1082635/how-to-decode-this-seemingly-gbk-encoded-string/1082640
func decode(encoded string) (string, error) {

	dec := new(mime.WordDecoder)

	var result []string
	arr := strings.Split(encoded, "\n")
	for _, src := range arr {
		if strings.HasPrefix(src, "=?") == false || strings.HasSuffix(src, "?=") == false {
			// add head "=?utf-8?q?"
			src = "=?utf-8?q?" + src
			// and replace tail "=" with "?="
			if strings.HasSuffix(src, "=") {
				src = src[:len(src)-1] + "?="
			} else {
				src = src + "?="
			}
		}
		str, err := dec.Decode(src)
		if err != nil {
			return encoded, err
		}
		result = append(result, str)
	}
	return strings.Join(result, ""), nil
}
