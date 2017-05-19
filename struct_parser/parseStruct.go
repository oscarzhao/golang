package structparser

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

// parse time for convenient
var (
	timeType    = reflect.TypeOf(time.Time{})
	timePtrType = reflect.TypeOf(&time.Time{})
)

// ParserFunc defines the signature of a parser
type ParserFunc func(reflect.Type) (*Field, error)

// Field is used to store the information of a struct field
type Field struct {
	Name        string
	Kind        string
	Type        string
	JSONName    string
	Required    bool
	Description string

	// if the Field is a struct, store all its fields in the list
	// if the Field is an array, store its field in the first element of the list
	Fields []Field
}

// Parse a type
func Parse(typ reflect.Type) (*Field, error) {
	originTyp := typ
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	parser, ok := getParser(typ)
	if !ok {
		fmt.Printf("Parse, cannot find parser; got %s", originTyp)
		return nil, fmt.Errorf("cannot find parser func; got %s", originTyp)
	}
	return parser(typ)
}

func getParser(typ reflect.Type) (ParserFunc, bool) {
	kind := typ.Kind()
	if kind == reflect.Ptr {
		kind = typ.Elem().Kind()
	}
	switch kind {
	case reflect.Bool:
		return parseSimpleType, true
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return parseSimpleType, true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return parseSimpleType, true
	case reflect.Float32, reflect.Float64:
		return parseSimpleType, true
	case reflect.String:
		return parseSimpleType, true
	case reflect.Array, reflect.Slice:
		return parseArrayType, true
	case reflect.Struct:
		return parseStructType, true
	default:
		return nil, false
	}
}

// parseSimpleType parses a certain pointer value and return its type
// pointer is ignored
func parseSimpleType(typ reflect.Type) (*Field, error) {
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	if typ.Kind() > reflect.Float64 && typ.Kind() != reflect.String {
		return nil, fmt.Errorf("dest must be a (ptr to) simple type; got %s", typ)
	}
	info := Field{Kind: typ.Kind().String(), Type: typ.String()}
	return &info, nil
}

// parseArrayType parses a certain pointer value and return its type
// pointer is ignored
func parseArrayType(typ reflect.Type) (*Field, error) {
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Array && typ.Kind() != reflect.Slice {
		return nil, fmt.Errorf("dest must be a (ptr to) array/slice type; got %s", typ)
	}
	info := Field{Kind: typ.Kind().String(), Type: typ.String()}
	elemType := typ.Elem()
	parser, ok := getParser(elemType)
	if !ok {
		return nil, fmt.Errorf("cannot find parser for field; got %s", elemType)
	}
	subInfo, err := parser(elemType)
	if err != nil {
		return nil, err
	}
	info.Fields = append(info.Fields, *subInfo)
	return &info, nil
}

// parseStructType parses a certain pointer value and return its swagger info
// pointer is ignored
func parseStructType(typ reflect.Type) (*Field, error) {
	originTyp := typ
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct {
		return nil, fmt.Errorf("dest must be a (ptr to) array/slice type; got %s", originTyp.Kind())
	}
	var infos []Field
	fieldCount := typ.NumField()
	for i := 0; i < fieldCount; i++ {
		field := typ.Field(i)

		var info Field
		info.Name = field.Name
		info.JSONName = field.Tag.Get("json")

		if field.Type.Kind() == reflect.Ptr {
			info.Kind = field.Type.Elem().Kind().String()
		} else {
			info.Kind = field.Type.Kind().String()
		}

		if field.Type == originTyp {
			return nil, fmt.Errorf("found nested type declaration, type name: %s", originTyp)
		}

		// calculate swagger info
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

		// filter out time.Time
		if field.Type == timePtrType || field.Type == timeType {
			info.Type = timeType.String()
			infos = append(infos, info)
			continue
		}

		// parse type information and nested structure
		parser, ok := getParser(field.Type)
		if !ok {
			return nil, fmt.Errorf("cannot find parser for field %s; got %s", field.Name, field.Type)
		}
		typeInfo, err := parser(field.Type)
		if err != nil {
			fmt.Printf("parse field %s failed, type:%s, err:%s\n", field.Name, field.Type, err)
			return nil, err
		}
		info.Type = typeInfo.Type
		info.Fields = typeInfo.Fields

		infos = append(infos, info)
	}
	info := &Field{
		Type:   typ.Name(),
		Fields: infos,
	}
	return info, nil
}
