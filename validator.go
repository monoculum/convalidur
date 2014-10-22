package govalidator

import (
	"errors"
	"regexp"
	"fmt"
	"strconv"
	"time"
)

var (
	ErrRequired = errors.New("the value is required")
	ErrNotMatched = errors.New("the value is not a valid email")
	ErrNotEqual = errors.New("the value is not a equal to value passed")
	ErrOutRange = errors.New("the value is out of range")
	ErrIn = errors.New("the value is not contains in values passed")
	ErrLen = errors.New("the value is more length than value passed")
	ErrDate = errors.New("the value is not a valid datetime")
)

const (
	PATTERN_EMAIL = `pattern_email`
	PATTERN_URL = `pattern_url`
)

type Validator struct {
	Errors *map[string][]string
}

func New(errors *map[string][]string) *Validator {
	return &Validator{errors}
}

type Str struct {
	validator *Validator
	value     string
	field     string
}

func (v *Validator) Str(value, field string) *Str {
	return &Str{v, value, field}
}

func (str *Str) Required() *Str {
	if str.value == "" {
		(*str.validator.Errors)[str.field] = append((*str.validator.Errors)[str.field], ErrRequired.Error())
	}
	return str
}

func (str *Str) Equal(eq string) *Str {
	if str.value == "" {
		return str
	}
	if str.value != eq {
		(*str.validator.Errors)[str.field] = append((*str.validator.Errors)[str.field], ErrNotEqual.Error())
	}
	return str
}

func (str *Str) Len(int int) *Str {
	if str.value == "" {
		return str
	}
	if len(str.value) != int {
		(*str.validator.Errors)[str.field] = append((*str.validator.Errors)[str.field], ErrLen.Error())
	}
	return str
}

func (str *Str) Range(min, max int) *Str {
	if str.value == "" {
		return str
	}
	if len(str.value) < min || len(str.value) > max {
		(*str.validator.Errors)[str.field] = append((*str.validator.Errors)[str.field], ErrOutRange.Error())
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
	(*str.validator.Errors)[str.field] = append((*str.validator.Errors)[str.field], ErrIn.Error())
	return str
}

func (str *Str) Date(layout string) *Str {
	if str.value == "" {
		return str
	}
	if _, err := time.Parse(layout, str.value); err != nil {
		(*str.validator.Errors)[str.field] = append((*str.validator.Errors)[str.field], ErrDate.Error())
	}
	return str
}

func (str *Str) Email() *Str {
	if str.value == "" {
		return str
	}
	return str.Regexp(PATTERN_EMAIL)
}

func (str *Str) URL() *Str {
	if str.value == "" {
		return str
	}
	return str.Regexp(PATTERN_URL)
}

func (str *Str) Regexp(pattern string) *Str {
	if str.value == "" {
		return str
	}
	if matched, err := regexp.Match(pattern, []byte(str.value)); err != nil {
		fmt.Printf("govalidator: %v", err)
	} else if !matched {
		(*str.validator.Errors)[str.field] = append((*str.validator.Errors)[str.field], ErrNotMatched.Error())
	}
	return str
}

type Int struct {
	validator *Validator
	value     int
	field     string
}

func (v *Validator) Int(value int, field string) *Int {
	return &Int{v, value, field}
}

func (int *Int) Required() *Int {
	if int.value == 0 {
		(*int.validator.Errors)[int.field] = append((*int.validator.Errors)[int.field], ErrRequired.Error())
	}
	return int
}

func (int *Int) Len(num int) *Int {
	if int.value == 0 {
		return int
	}
	if len(strconv.Itoa(int.value)) != num {
		(*int.validator.Errors)[int.field] = append((*int.validator.Errors)[int.field], ErrLen.Error())
	}
	return int
}

func (int *Int) Equal(eq int) *Int {
	if int.value == 0 {
		return int
	}
	if int.value != eq {
		(*int.validator.Errors)[int.field] = append((*int.validator.Errors)[int.field], ErrNotEqual.Error())
	}
	return int
}

func (int *Int) Range(min, max int) *Int {
	if int.value == 0 {
		return int
	}
	if int.value < min || int.value > max {
		(*int.validator.Errors)[int.field] = append((*int.validator.Errors)[int.field], ErrOutRange.Error())
	}
	return int
}

type Slice struct {
	validator *Validator
	value     interface {}
	field     string
}

func (v *Validator) Slice(value interface {}, field string) *Slice {
	return &Slice{v, value, field}
}

func (sl *Slice) Required() *Slice {
	return sl
}

func (sl *Slice) Range(min, max int) *Slice {
	return sl
}

func (sl *Slice) Contains(values interface {}) *Slice {
	return sl
}

type Map struct {
	validator *Validator
	value     interface {}
	field     string
}

func (v *Validator) Map(value interface {}, field string) *Map {
	return &Map{v, value, field}
}

func (sl *Map) Required() *Map {
	return sl
}

func (sl *Map) Range(min, max int) *Map {
	return sl
}

func (sl *Map) Contains(values interface {}) *Map {
	return sl
}

func (sl *Map) Date(layout string) *Map {
	return sl
}
