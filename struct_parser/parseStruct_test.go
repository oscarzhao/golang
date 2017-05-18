package structparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseStructTag(t *testing.T) {
	// Person ...
	type Person struct {
		ID      int64   `json:"id" swagger:"Required, ID of the person"`
		Name    string  `json:"name" swagger:"Required, Name of the person"`
		Age     int     `json:"age" swagger:"Optional, Age of the person"`
		Hobbies *string `json:"hobbies" swagger:"Optional, Hobbies of the person"`
	}
	p := &Person{
		ID:   1,
		Name: "oscar",
		Age:  22,
	}
	fields, err := ParseStructInfo(p)
	assert.Nil(t, err, "parse Person struct should succeed")
	assert.Equal(t, 4, len(fields), "check struct Person's field number")
	assert.Equal(t, "int64", fields[0].Type)
	assert.Equal(t, "string", fields[1].Type)
	assert.Equal(t, "int", fields[2].Type)
	assert.Equal(t, "string", fields[3].Type)
}
