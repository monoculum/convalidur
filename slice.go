package convalidur

import "reflect"

type Slice struct {
	raw    interface{}
	field  string
	errors *map[string][]Err

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
			(*sl.errors)[sl.field] = append((*sl.errors)[sl.field], Err{ErrRequired, CodeRequired})
		}
	default:
		(*sl.errors)[sl.field] = append((*sl.errors)[sl.field], Err{ErrUnsupported, CodeUnsupported})
	}
	return sl
}

func (sl *Slice) Range(min, max int) *Slice {
	sl.value = reflect.ValueOf(sl.raw)
	if sl.value.Kind() == reflect.Ptr {
		sl.value = sl.value.Elem()
	}
	if !sl.value.IsNil() {
		switch sl.value.Kind() {
		case reflect.Slice, reflect.Array:
			len := sl.value.Len()
			if len < min || len > max {
				(*sl.errors)[sl.field] = append((*sl.errors)[sl.field], Err{ErrOutRange, CodeOutRange})
			}
		default:
			(*sl.errors)[sl.field] = append((*sl.errors)[sl.field], Err{ErrUnsupported, CodeUnsupported})
		}
	}
	return sl
}

func (sl *Slice) In(values interface{}) *Slice {
	sl.value = reflect.ValueOf(sl.raw)
	if sl.value.Kind() == reflect.Ptr {
		sl.value = sl.value.Elem()
	}
	if !sl.value.IsNil() {
		switch sl.value.Kind() {
		case reflect.Slice, reflect.Array:
			value := sl.value
			len := value.Len()
			for i := 0; i < len; i++ {
				sl.value = value.Index(i)
				sl.in(values)
			}
		default:
			(*sl.errors)[sl.field] = append((*sl.errors)[sl.field], Err{ErrUnsupported, CodeUnsupported})
		}
	}
	return sl
}

func (sl *Slice) Len(le int) *Slice {
	sl.value = reflect.ValueOf(sl.raw)
	if sl.value.Kind() == reflect.Ptr {
		sl.value = sl.value.Elem()
	}
	if !sl.value.IsNil() {
		switch sl.value.Kind() {
		case reflect.Slice, reflect.Array:
			if sl.value.Len() != le {
				(*sl.errors)[sl.field] = append((*sl.errors)[sl.field], Err{ErrLen, CodeLen})
			}
		default:
			(*sl.errors)[sl.field] = append((*sl.errors)[sl.field], Err{ErrUnsupported, CodeUnsupported})
		}
	}
	return sl
}

func (sl *Slice) in(n interface{}) *Slice {
	found := false

	switch sl.value.Kind() {
	case reflect.Slice, reflect.Array:
		value := sl.value
		len := value.Len()
		for i := 0; i < len; i++ {
			sl.value = value.Index(i)
			sl.in(n)
		}
	case reflect.Map:
		value := sl.value
		keys := sl.value.MapKeys()
	for _, key := range keys {
		sl.value = value.MapIndex(key)
		sl.in(n)
	}
	case reflect.String:
		values := reflect.ValueOf(n)
		switch values.Kind() {
		case reflect.Slice, reflect.Array:
			for j := 0; j < values.Len(); j++ {
				value := values.Index(j)
				str := ""
				switch value.Kind() {
				case reflect.String:
					str = value.String()
				case reflect.Interface:
					if m, ok := value.Interface().(string); ok {
						str = m
					} else if value.Kind() == reflect.String {
						str = value.String()
					}
				default:
					(*sl.errors)[sl.field] = append((*sl.errors)[sl.field], Err{ErrBadParameter, CodeBadParameter})
				}
				if sl.value.String() == str {
					found = true
				}
			}
		default:
			(*sl.errors)[sl.field] = append((*sl.errors)[sl.field], Err{ErrBadParameter, CodeBadParameter})
		}
		if !found {
			(*sl.errors)[sl.field] = append((*sl.errors)[sl.field], Err{ErrIn, CodeIn})
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		values := reflect.ValueOf(n)
		switch values.Kind() {
		case reflect.Slice, reflect.Array:
			for j := 0; j < values.Len(); j++ {
				value := values.Index(j)
				num := 0
				switch value.Kind() {
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					num = int(value.Int())
				case reflect.Interface:
					switch p := value.Interface().(type) {
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
						switch value.Kind() {
						case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
							num = int(value.Int())
						default:
							(*sl.errors)[sl.field] = append((*sl.errors)[sl.field], Err{ErrBadParameter, CodeBadParameter})
						}
					}
				default:
					(*sl.errors)[sl.field] = append((*sl.errors)[sl.field], Err{ErrBadParameter, CodeBadParameter})
				}
				if int(sl.value.Int()) == num {
					found = true
				}
			}
		default:
			(*sl.errors)[sl.field] = append((*sl.errors)[sl.field], Err{ErrBadParameter, CodeBadParameter})
		}
		if !found {
			(*sl.errors)[sl.field] = append((*sl.errors)[sl.field], Err{ErrIn, CodeIn})
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
		(*sl.errors)[sl.field] = append((*sl.errors)[sl.field], Err{ErrUnsupported, CodeUnsupported})
	}

	return sl
}
