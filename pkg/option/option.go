package option

import (
	"errors"
	"net/url"
	"reflect"
)

var (
	ErrRequiredMissing = errors.New("required field missing")
)

type Option interface {
	// Encode encodes '*Options' as 'URL encoded' format.
	//
	// Since the options for each API will be different,
	// we want to implement own Encode() method to APIs.
	// All options must have JSON tag.
	Encode() (string, error)
}

// EncodeQueryParam parses the `option` structure.
func EncodeQueryParam(opt interface{}) string {
	response := url.Values{}
	for k, v := range structToMap(opt) {
		if v == "" || v == false {
			continue
		}

		var value string

		// Couldn't make: ?myBool="false"
		switch reflect.TypeOf(v).Kind() {
		case reflect.String:
			value = v.(string)
		case reflect.Bool:
			value = "true"
		default:
			panic("unsupport option type")
		}

		response.Set(k, value)
	}

	return response.Encode()
}

func structToMap(opt interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	val := reflect.ValueOf(opt)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		tag := typ.Field(i).Tag.Get("json")
		if tag == "" {
			panic("must set tag to option")
		}
		var value interface{}
		if field.Kind() == reflect.Struct {
			value = structToMap(field.Interface())
		} else {
			value = field.Interface()
		}
		m[tag] = value
	}
	return m
}
