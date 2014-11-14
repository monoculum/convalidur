package valider

import (
	"reflect"
	"time"
)

type Map struct {
	raw    interface{}
	field  string
	errors *map[string][]Err

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
			(*ma.errors)[ma.field] = append((*ma.errors)[ma.field], Err{ErrRequired, CodeRequired})
		}
	default:
		(*ma.errors)[ma.field] = append((*ma.errors)[ma.field], Err{ErrUnsupported, CodeUnsupported})
	}
	return ma
}

func (ma *Map) Keys(keys ...string) *Map {
	ma.value = reflect.ValueOf(ma.raw)
	if ma.value.Kind() == reflect.Ptr {
		ma.value = ma.value.Elem()
	}
	if ma.value.Len() != 0 {
		switch ma.value.Kind() {
		case reflect.Map:
		for _, key := range keys {
			if !ma.value.MapIndex(reflect.ValueOf(key)).IsValid() {
				(*ma.errors)[ma.field] = append((*ma.errors)[ma.field], Err{ErrNotFound, CodeNotFound})
			}
		}
		default:
			(*ma.errors)[ma.field] = append((*ma.errors)[ma.field], Err{ErrUnsupported, CodeUnsupported})
		}
	}
	return ma
}

func (ma *Map) Range(min, max int) *Map {
	ma.value = reflect.ValueOf(ma.raw)
	if ma.value.Kind() == reflect.Ptr {
		ma.value = ma.value.Elem()
	}
	if ma.value.Len() != 0 {
		switch ma.value.Kind() {
		case reflect.Map:
			len := ma.value.Len()
			if len < min || len > max {
				(*ma.errors)[ma.field] = append((*ma.errors)[ma.field], Err{ErrOutRange, CodeOutRange})
			}
		default:
			(*ma.errors)[ma.field] = append((*ma.errors)[ma.field], Err{ErrUnsupported, CodeUnsupported})
		}
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
	if ma.value.Len() != 0 {
		switch ma.value.Kind() {
		case reflect.Map:
		for _, key := range ma.value.MapKeys() {
			ma.value = ma.value.MapIndex(key)
			ma.date(layout)
		}
		default:
			(*ma.errors)[ma.field] = append((*ma.errors)[ma.field], Err{ErrUnsupported, CodeUnsupported})
		}
	}
	return ma
}

func (ma *Map) date(layout string) *Map {
	switch ma.value.Kind() {
	case reflect.Slice, reflect.Array:
		len := ma.value.Len()
		for i := 0; i < len; i++ {
			ma.value = ma.value.Index(i)
			ma.date(layout)
		}
	case reflect.String:
		if _, err := time.Parse(layout, ma.value.String()); err != nil {
			(*ma.errors)[ma.field] = append((*ma.errors)[ma.field], Err{ErrDate, CodeDate})
		}
	default:
		(*ma.errors)[ma.field] = append((*ma.errors)[ma.field], Err{ErrUnsupported, CodeUnsupported})
	}
	return ma
}
