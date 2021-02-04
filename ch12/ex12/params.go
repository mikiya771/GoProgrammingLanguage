package params

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

var validators map[string]func(string) bool

func init() {
	validators = map[string]func(string) bool{}
	validators["zip"] = func(name string) bool {
		//xxx-xxxx is valid
		twoWords := strings.Split(name, "-")
		if len(twoWords) != 2 {
			return false
		}
		num, err := strconv.Atoi(twoWords[0])
		if err != nil {
			return false
		}
		if num > 999 || num < 100 {
			return false
		}

		num, err = strconv.Atoi(twoWords[1])
		if err != nil {
			return false
		}
		if num > 9999 || num < 1000 {
			return false
		}
		return true
	}
}

func unpack(urlVals url.Values, ptr interface{}) error {
	fields := make(map[string]reflect.Value)
	v := reflect.ValueOf(ptr).Elem()
	localValidators := map[string][]string{}
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		tag := fieldInfo.Tag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		functionNames := tag.Get("valdator")
		funcNameArr := strings.Split(functionNames, ",")
		localValidators[name] = funcNameArr
		fields[name] = v.Field(i)
	}

	for name, values := range urlVals {
		f := fields[name]
		if !f.IsValid() {
			continue
		}
		for _, value := range values {
			for _, fn := range localValidators[name] {
				if !validators[fn](value) {
					return fmt.Errorf("validate[%s](%s) is false", fn, value)
				}
			}
			if f.Kind() == reflect.Slice {
				elem := reflect.New(f.Type().Elem()).Elem()
				if err := populate(elem, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
				f.Set(reflect.Append(f, elem))
			} else {
				if err := populate(f, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
			}
		}
	}
	return nil
}
func Unpack(req *http.Request, ptr interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}
	return unpack(req.Form, ptr)
}

//!-Unpack

//!+populate
func populate(v reflect.Value, value string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)

	case reflect.Int:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(i)

	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		v.SetBool(b)

	default:
		return fmt.Errorf("unsupported kind %s", v.Type())
	}
	return nil
}

//!-populate
