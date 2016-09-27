/******************************************************************************
通过 base64 转码 []byte 格式的 uuid
******************************************************************************/
package encoding

import (
	"encoding/base64"
	"fmt"

	"github.com/satori/go.uuid"
)

func main() {
	input := uuid.NewV4().Bytes()
	out := base64.StdEncoding.EncodeToString(input)
	fmt.Printf("input: %v, len:%d\n", input, len(input)) // expect [input: xxx, len:16]
	fmt.Printf("out: %s, len:%d\n", out, len(out))       // expect [out: xxx, len:24]
}
