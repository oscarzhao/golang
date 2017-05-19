package structparser

import (
	"testing"
	"time"

	"reflect"

	"github.com/stretchr/testify/assert"
)

func TestParseStructTag(t *testing.T) {
	type Crime struct {
		Time     time.Time `json:"time" swagger:"Required, the time crime is committed"`
		Place    string    `json:"place" swagger:"Required, the place where crime is committed"`
		Severity int8      `json:"severity" swagger:"Required, the severity of the crime"`
	}
	// Person ...
	type Person struct {
		ID           int64     `json:"id" swagger:"Required, ID of the person"`
		Name         string    `json:"name" swagger:"Required, Name of the person"`
		Age          int       `json:"age" swagger:"Optional, Age of the person"`
		Hobbies      *string   `json:"hobbies" swagger:"Optional, Hobbies of the person"`
		CrimeHistory []Crime   `json:"crimeHistory" swagger:"Optional, Crime history of the person"`
		Created      time.Time `json:"created" swagger:"Required, Created time of the record"`
	}
	p := &Person{}
	field, err := Parse(reflect.TypeOf(p))
	assert.Nil(t, err, "parse Person struct should succeed")
	assert.Equal(t, 6, len(field.Fields), "check struct Person's field number")

	assert.Equal(t, "int64", field.Fields[0].Kind, "test ID")
	assert.Equal(t, "int64", field.Fields[0].Type, "test ID")
	assert.Equal(t, "id", field.Fields[0].JSONName, "test ID")

	assert.Equal(t, "string", field.Fields[1].Kind, "test Name")
	assert.Equal(t, "string", field.Fields[1].Type, "test Name")
	assert.Equal(t, "name", field.Fields[1].JSONName, "test Name")

	assert.Equal(t, "int", field.Fields[2].Kind, "test Age")
	assert.Equal(t, "int", field.Fields[2].Type, "test Age")
	assert.Equal(t, "age", field.Fields[2].JSONName, "test Age")

	assert.Equal(t, "string", field.Fields[3].Kind, "test Hobbies")
	assert.Equal(t, "string", field.Fields[3].Type, "test Hobbies")
	assert.Equal(t, "hobbies", field.Fields[3].JSONName, "test Hobbies")

	assert.Equal(t, "slice", field.Fields[4].Kind, "test CrimeHistory")
	assert.Equal(t, "[]structparser.Crime", field.Fields[4].Type, "test CrimeHistory")
	assert.Equal(t, "crimeHistory", field.Fields[4].JSONName, "test CrimeHistory")

	crimeHistoryList := field.Fields[4]
	assert.Equal(t, 1, len(crimeHistoryList.Fields), "test field count of []Crime")

	crimeHistoryInfo := crimeHistoryList.Fields[0]

	assert.Equal(t, 3, len(crimeHistoryInfo.Fields), "test fields count of Crime")
	assert.Equal(t, "int8", crimeHistoryInfo.Fields[2].Kind)
	assert.Equal(t, "int8", crimeHistoryInfo.Fields[2].Type)
	assert.Equal(t, "severity", crimeHistoryInfo.Fields[2].JSONName)

	assert.Equal(t, "struct", field.Fields[5].Kind, "test Created")
	assert.Equal(t, "time.Time", field.Fields[5].Type, "test Created")
	assert.Equal(t, "created", field.Fields[5].JSONName, "test Created")
}
