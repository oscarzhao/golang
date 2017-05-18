package structparser

import (
	"fmt"
	"reflect"
	"strings"
)

// Field is used to store the information of a struct field
type Field struct {
	Name        string
	Type        string
	JSONName    string
	Required    bool
	Description string
}

// ParseStructInfo parses a certain pointer value and return its swagger info
func ParseStructInfo(ifc interface{}) ([]Field, error) {
	typ := reflect.TypeOf(ifc)
	if typ.Kind() != reflect.Ptr || typ.Elem().Kind() != reflect.Struct {
		return nil, fmt.Errorf("dest must be pointer to struct; got %T", typ)
	}
	var infos []Field
	elem := typ.Elem()
	fieldCount := elem.NumField()
	for i := 0; i < fieldCount; i++ {
		field := elem.Field(i)

		var info Field
		info.Name = field.Name
		info.Type = field.Type.String()
		info.JSONName = field.Tag.Get("json")

		swaggerTag := field.Tag.Get("swagger")
		items := strings.Split(swaggerTag, ",")
		if len(items) < 2 {
			return nil, fmt.Errorf("field:%s, invalid swagger tag: %s", field.Name, swaggerTag)
		}
		if requireStr := strings.Trim(items[0], " "); requireStr == "Required" {
			info.Required = true
		} else if requireStr == "Optional" {
			info.Required = false
		} else {
			return nil, fmt.Errorf("field:%s, invalid swagger tag: %s", field.Name, swaggerTag)
		}
		info.Description = strings.Trim(items[1], " ")
		infos = append(infos, info)
	}
	return infos, nil
}
