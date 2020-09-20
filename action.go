package terrapi

import (
	"reflect"
)

type Action interface {
	unmarshal() ([]string, error)
	setOutErr([]byte, []byte)
	GetOutErr() ([]byte, []byte)
}

func unmarshalAction(a Action) ([]string, error) {
	v := reflect.ValueOf(a)
	if v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil, &InvalidActionError{Type: v.Type()}
	}

	c := v.NumField()

	b := []string{v.Field(0).String()}

	// we count from 1 to bypass action name
	for i := 1; i < c; i++ {
		fieldValue := v.Field(i)
		fieldType := v.Type().Field(i)

		switch fieldValue.Kind() {
		case reflect.Bool:
			if fieldValue.Bool() {
				b = append(b, fieldType.Tag.Get("cli"))
			}
		case reflect.String:
			if !fieldValue.IsZero() {
				b = append(b, fieldType.Tag.Get("cli") + fieldValue.String())
			}
		case reflect.Slice:
			s, ok := fieldValue.Interface().([]string)
			if ok {
				for _, j := range s {
					if j != "" {
						b = append(b, fieldType.Tag.Get("cli") + j)
					}
				}
			}
		}
	}

	return b, nil
}
