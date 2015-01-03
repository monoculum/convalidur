package valider

import (
	"errors"
)

var (
	ErrRequired    = errors.New("is required")
	ErrNotMatched  = errors.New("not matched")
	ErrNotEqual    = errors.New("is not equal to value passed")
	ErrEqual       = errors.New("is equal to value passed")
	ErrOutRange    = errors.New("is out of range")
	ErrIn          = errors.New("is not in the values passed")
	ErrLen         = errors.New("is more length than value passed")
	ErrDate        = errors.New("is not a valid datetime")
	ErrNotFound    = errors.New("not found the value")
	ErrExists      = errors.New("the value exists")

	ErrUnsupported  = errors.New("unsopported type")
	ErrBadParameter = errors.New("bad parameter")
)

const (
	CodeRequired    = "required"
	CodeNotMatched  = "not_matched"
	CodeNotEqual    = "not_equal"
	CodeEqual       = "equal"
	CodeOutRange    = "out_range"
	CodeIn          = "in"
	CodeLen         = "length"
	CodeDate        = "not_date"
	CodeNotFound    = "not_found"
	CodeExists      = "exists"

	CodeUnsupported  = "unsopported_type"
	CodeBadParameter = "bad_parameter"
)

const (
	PATTERN_EMAIL = "[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?"
	PATTERN_URL   = `^((ftp|http|https):\/\/)?(\S+(:\S*)?@)?((([1-9]\d?|1\d\d|2[01]\d|22[0-3])(\.(1?\d{1,2}|2[0-4]\d|25[0-5])){2}(?:\.([0-9]\d?|1\d\d|2[0-4]\d|25[0-4]))|((www\.)?)?(([a-z\x{00a1}-\x{ffff}0-9]+-?-?_?)*[a-z\x{00a1}-\x{ffff}0-9]+)(?:\.([a-z\x{00a1}-\x{ffff}]{2,}))?)|localhost)(:(\d{1,5}))?((\/|\?|#)[^\s]*)?$`
)

type Errors map[string][]Error

func (f *Errors) Add(name string, err error, code string) {
	(*f)[name] = append((*f)[name], Error{err, code})
}

type Error struct {
	Error  error
	Code   string
}

type Validator struct {
	Errors *Errors
}

func New(errors *Errors) *Validator {
	return &Validator{errors}
}
