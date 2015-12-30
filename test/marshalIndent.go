package main
/**
 * marshal a struct into a html displayable json string
 */
import (
	"encoding/json"
	"log"
	"fmt"
)

func main() {
	type Road struct {
		Name   string
		Number int
	}
	roads := []Road{
		{"Diamond Fork", 29},
		{"Sheep Creek", 51},
	}

	b, err := json.MarshalIndent(roads, "<br>", "&nbsp;&nbsp;&nbsp;&nbsp;")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
