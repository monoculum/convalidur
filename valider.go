package valider

import (
	"errors"
	"reflect"
	"regexp"
	"strconv"
	"time"
)

var (
	ErrRequired   = errors.New("is required")
	ErrNotMatched = errors.New("is not a valid email")
	ErrNotEqual   = errors.New("is not a equal to value passed")
	ErrOutRange   = errors.New("is out of range")
	ErrIn         = errors.New("is not in the values passed")
	ErrLen        = errors.New("is more length than value passed")
	ErrDate       = errors.New("is not a valid datetime")

	ErrUnsupported  = errors.New("unsopported type")
	ErrBadParameter = errors.New("bad parameter")
)

const (
	PATTERN_EMAIL = "[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?"
	PATTERN_URL   = `(?i)^([a-z]([a-z]|\d|\+|-|\.)*):(\/\/(((([a-z]|\d|-|\.|_|~|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])|(%[\da-f]{2})|[!\$&'\(\)\*\+,;=]|:)*@)?((\[(|(v[\da-f]{1,}\.(([a-z]|\d|-|\.|_|~)|[!\$&'\(\)\*\+,;=]|:)+))\])|((\d|[1-9]\d|1\d\d|2[0-4]\d|25[0-5])\.(\d|[1-9]\d|1\d\d|2[0-4]\d|25[0-5])\.(\d|[1-9]\d|1\d\d|2[0-4]\d|25[0-5])\.(\d|[1-9]\d|1\d\d|2[0-4]\d|25[0-5]))|(([a-z]|\d|-|\.|_|~|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])|(%[\da-f]{2})|[!\$&'\(\)\*\+,;=])*)(:\d*)?)(\/(([a-z]|\d|-|\.|_|~|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])|(%[\da-f]{2})|[!\$&'\(\)\*\+,;=]|:|@)*)*|(\/((([a-z]|\d|-|\.|_|~|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])|(%[\da-f]{2})|[!\$&'\(\)\*\+,;=]|:|@)+(\/(([a-z]|\d|-|\.|_|~|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])|(%[\da-f]{2})|[!\$&'\(\)\*\+,;=]|:|@)*)*)?)|((([a-z]|\d|-|\.|_|~|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])|(%[\da-f]{2})|[!\$&'\(\)\*\+,;=]|:|@)+(\/(([a-z]|\d|-|\.|_|~|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])|(%[\da-f]{2})|[!\$&'\(\)\*\+,;=]|:|@)*)*)|((([a-z]|\d|-|\.|_|~|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])|(%[\da-f]{2})|[!\$&'\(\)\*\+,;=]|:|@)){0})(\?((([a-z]|\d|-|\.|_|~|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])|(%[\da-f]{2})|[!\$&'\(\)\*\+,;=]|:|@)|[\uE000-\uF8FF]|\/|\?)*)?(\#((([a-z]|\d|-|\.|_|~|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])|(%[\da-f]{2})|[!\$&'\(\)\*\+,;=]|:|@)|\/|\?)*)?$`
)

type Validator struct {
	Errors *map[string][]string
}

func New(errors *map[string][]string) *Validator {
	return &Validator{errors}
}

type Str struct {
	value  string
	field  string
	errors *map[string][]string
}

func (v *Validator) Str(value, field string) *Str {
	return &Str{value, field, v.Errors}
}

func (str *Str) Required() *Str {
	if str.value == "" {
		(*str.errors)[str.field] = append((*str.errors)[str.field], ErrRequired.Error())
	}
	return str
}

func (str *Str) Equal(eq string) *Str {
	if str.value != "" && str.value != eq {
		(*str.errors)[str.field] = append((*str.errors)[str.field], ErrNotEqual.Error())
	}
	return str
}

func (str *Str) Len(int int) *Str {
	if str.value != "" && len(str.value) != int {
		(*str.errors)[str.field] = append((*str.errors)[str.field], ErrLen.Error())
	}
	return str
}

func (str *Str) Range(min, max int) *Str {
	if str.value != "" && len(str.value) < min || len(str.value) > max {
		(*str.errors)[str.field] = append((*str.errors)[str.field], ErrOutRange.Error())
	}
	return str
}

func (str *Str) In(values ...string) *Str {
	if str.value == "" {
		return str
	}
	for _, value := range values {
		if str.value == value {
			return str
		}
	}
	(*str.errors)[str.field] = append((*str.errors)[str.field], ErrIn.Error())
	return str
}

func (str *Str) Date(layout string) *Str {
	if str.value == "" {
		return str
	}
	if _, err := time.Parse(layout, str.value); err != nil {
		(*str.errors)[str.field] = append((*str.errors)[str.field], ErrDate.Error())
	}
	return str
}

func (str *Str) Email() *Str {
	if str.value == "" {
		return str
	}
	return str.RegExp(PATTERN_EMAIL)
}

func (str *Str) URL() *Str {
	if str.value == "" {
		return str
	}
	return str.RegExp(PATTERN_URL)
}

func (str *Str) RegExp(pattern string) *Str {
	if str.value == "" {
		return str
	}
	if matched, err := regexp.MatchString(pattern, str.value); err != nil {
		(*str.errors)[str.field] = append((*str.errors)[str.field], ErrBadParameter.Error())
	} else if !matched {
		(*str.errors)[str.field] = append((*str.errors)[str.field], ErrNotMatched.Error())
	}
	return str
}

type Int struct {
	value  int
	field  string
	errors *map[string][]string
}

func (v *Validator) Int(value int, field string) *Int {
	return &Int{value, field, v.Errors}
}

func (int *Int) Required() *Int {
	if int.value == 0 {
		(*int.errors)[int.field] = append((*int.errors)[int.field], ErrRequired.Error())
	}
	return int
}

func (int *Int) Len(num int) *Int {
	if int.value != 0 && len(strconv.Itoa(int.value)) != num {
		(*int.errors)[int.field] = append((*int.errors)[int.field], ErrLen.Error())
	}
	return int
}

func (int *Int) Equal(eq int) *Int {
	if int.value != 0 && int.value != eq {
		(*int.errors)[int.field] = append((*int.errors)[int.field], ErrNotEqual.Error())
	}
	return int
}

func (int *Int) Range(min, max int) *Int {
	if int.value != 0 && (int.value < min || int.value > max) {
		(*int.errors)[int.field] = append((*int.errors)[int.field], ErrOutRange.Error())
	}
	return int
}

type Slice struct {
	raw    interface{}
	field  string
	errors *map[string][]string

	value reflect.Value
}

func (v *Validator) Slice(value interface{}, field string) *Slice {
	return &Slice{raw: value, field: field, errors: v.Errors}
}

func (sl *Slice) Required() *Slice {
	sl.value = reflect.ValueOf(sl.raw)
	if sl.value.Kind() == reflect.Ptr {
		sl.value = sl.value.Elem()
	}
	switch sl.value.Kind() {
	case reflect.Slice, reflect.Array:
		if sl.value.Len() == 0 {
			(*sl.errors)[sl.field] = append((*sl.errors)[sl.field], ErrRequired.Error())
		}
	default:
		(*sl.errors)[sl.field] = append((*sl.errors)[sl.field], ErrUnsupported.Error())
	}
	return sl
}

func (sl *Slice) Range(min, max int) *Slice {
	sl.value = reflect.ValueOf(sl.raw)
	if sl.value.Kind() == reflect.Ptr {
		sl.value = sl.value.Elem()
	}
	switch sl.value.Kind() {
	case reflect.Slice, reflect.Array:
		l := sl.value.Len()
		if l < min || l > max {
			(*sl.errors)[sl.field] = append((*sl.errors)[sl.field], ErrOutRange.Error())
		}
	default:
		(*sl.errors)[sl.field] = append((*sl.errors)[sl.field], ErrUnsupported.Error())
	}
	return sl
}

func (sl *Slice) In(values interface{}) *Slice {
	sl.value = reflect.ValueOf(sl.raw)
	if sl.value.Kind() == reflect.Ptr {
		sl.value = sl.value.Elem()
	}
	switch sl.value.Kind() {
	case reflect.Slice, reflect.Array:
		v := sl.value
		l := v.Len()
		for i := 0; i < l; i++ {
			sl.value = v.Index(i)
			sl.in(values)
		}
	default:
		(*sl.errors)[sl.field] = append((*sl.errors)[sl.field], ErrUnsupported.Error())
	}
	return sl
}

func (sl *Slice) in(n interface{}) *Slice {
	found := false

	switch sl.value.Kind() {
	case reflect.String:
		values := reflect.ValueOf(n)
		switch values.Kind() {
		case reflect.Slice, reflect.Array:
			for j := 0; j < values.Len(); j++ {
				t := values.Index(j)
				str := ""
				switch t.Kind() {
				case reflect.String:
					str = t.String()
				case reflect.Interface:
					if m, ok := t.Interface().(string); ok {
						str = m
					} else if t.Kind() == reflect.String {
						str = t.String()
					}
				default:
					(*sl.errors)[sl.field] = append((*sl.errors)[sl.field], ErrBadParameter.Error())
				}
				if sl.value.String() == str {
					found = true
				}
			}
		default:
			(*sl.errors)[sl.field] = append((*sl.errors)[sl.field], ErrBadParameter.Error())
		}
		if !found {
			(*sl.errors)[sl.field] = append((*sl.errors)[sl.field], ErrIn.Error())
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		values := reflect.ValueOf(n)
		switch values.Kind() {
		case reflect.Slice, reflect.Array:
			for j := 0; j < values.Len(); j++ {
				t := values.Index(j)
				num := 0
				switch t.Kind() {
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					num = int(t.Int())
				case reflect.Interface:
					switch p := t.Interface().(type) {
					case int:
						num = int(p)
					case int8:
						num = int(p)
					case int16:
						num = int(p)
					case int32:
						num = int(p)
					case int64:
						num = int(p)
					default:
						switch t.Kind() {
						case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
							num = int(t.Int())
						default:
							(*sl.errors)[sl.field] = append((*sl.errors)[sl.field], ErrBadParameter.Error())
						}
					}
				default:
					(*sl.errors)[sl.field] = append((*sl.errors)[sl.field], ErrBadParameter.Error())
				}
				if int(sl.value.Int()) == num {
					found = true
				}
			}
		default:
			(*sl.errors)[sl.field] = append((*sl.errors)[sl.field], ErrBadParameter.Error())
		}
		if !found {
			(*sl.errors)[sl.field] = append((*sl.errors)[sl.field], ErrIn.Error())
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		// TODO
	case reflect.Float32, reflect.Float64:
		// TODO
	case reflect.Bool:
		// TODO
	case reflect.Interface:
		// TODO
	default:
		(*sl.errors)[sl.field] = append((*sl.errors)[sl.field], ErrUnsupported.Error())
	}

	return sl
}

type Map struct {
	raw    interface{}
	field  string
	errors *map[string][]string

	value reflect.Value
}

func (v *Validator) Map(value interface{}, field string) *Map {
	return &Map{raw: value, field: field, errors: v.Errors}
}

func (ma *Map) Required() *Map {
	ma.value = reflect.ValueOf(ma.raw)
	if ma.value.Kind() == reflect.Ptr {
		ma.value = ma.value.Elem()
	}
	switch ma.value.Kind() {
	case reflect.Map:
		if ma.value.Len() == 0 {
			(*ma.errors)[ma.field] = append((*ma.errors)[ma.field], ErrRequired.Error())
		}
	default:
		(*ma.errors)[ma.field] = append((*ma.errors)[ma.field], ErrUnsupported.Error())
	}
	return ma
}

func (ma *Map) Range(min, max int) *Map {
	ma.value = reflect.ValueOf(ma.raw)
	if ma.value.Kind() == reflect.Ptr {
		ma.value = ma.value.Elem()
	}
	if ma.value.Len() == 0 {
		return ma
	}
	switch ma.value.Kind() {
	case reflect.Map:
		if ma.value.Len() < min || ma.value.Len() > max {
			(*ma.errors)[ma.field] = append((*ma.errors)[ma.field], ErrOutRange.Error())
		}
	default:
		(*ma.errors)[ma.field] = append((*ma.errors)[ma.field], ErrUnsupported.Error())
	}
	return ma
}

func (ma *Map) In(values interface{}) *Map {
	return ma
}

func (ma *Map) Date(layout string) *Map {
	ma.value = reflect.ValueOf(ma.raw)
	if ma.value.Kind() == reflect.Ptr {
		ma.value = ma.value.Elem()
	}
	if ma.value.Len() == 0 {
		return ma
	}
	switch ma.value.Kind() {
	case reflect.Map:
		for _, key := range ma.value.MapKeys() {
			ma.value = ma.value.MapIndex(key)
			ma.date(layout)
		}
	default:
		(*ma.errors)[ma.field] = append((*ma.errors)[ma.field], ErrUnsupported.Error())
	}
	return ma
}

func (ma *Map) date(layout string) *Map {
	switch ma.value.Kind() {
	case reflect.Slice, reflect.Array:
		l := ma.value.Len()
		for i := 0; i < l; i++ {
			ma.value = ma.value.Index(i)
			ma.date(layout)
		}
	case reflect.String:
		if _, err := time.Parse(layout, ma.value.String()); err != nil {
			(*ma.errors)[ma.field] = append((*ma.errors)[ma.field], ErrDate.Error())
		}
	default:
		(*ma.errors)[ma.field] = append((*ma.errors)[ma.field], ErrUnsupported.Error())
	}
	return ma
}
