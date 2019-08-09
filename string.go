package valider

import (
	"regexp"
	"time"
)

type Str struct {
	value  string
	field  string
	errors Errors
}

func (v *Validator) Str(value, field string) *Str {
	return &Str{value, field, v.Errors}
}

func (str *Str) Required() *Str {
	if str.value == "" {
		str.errors[str.field] = append(str.errors[str.field], Error{ErrRequired, CodeRequired, nil})
	}
	return str
}

func (str *Str) Equal(eq string) *Str {
	if str.value != "" && str.value != eq {
		str.errors[str.field] = append(str.errors[str.field], Error{ErrNotEqual, CodeNotEqual, eq})
	}
	return str
}

func (str *Str) NotEqual(eq string) *Str {
	if str.value != "" && str.value == eq {
		str.errors[str.field] = append(str.errors[str.field], Error{ErrNotEqual, CodeNotEqual, eq})
	}
	return str
}

func (str *Str) Len(num int) *Str {
	if str.value != "" && len(str.value) != num {
		str.errors[str.field] = append(str.errors[str.field], Error{ErrLen, CodeLen, num})
	}
	return str
}

func (str *Str) Range(min, max int) *Str {
	len := len(str.value)
	if str.value != "" && len < min || len > max {
		str.errors[str.field] = append(str.errors[str.field], Error{ErrOutRange, CodeOutRange, []int{min, max}})
	}
	return str
}

func (str *Str) In(values ...string) *Str {
	if str.value != "" {
		for _, value := range values {
			if str.value == value {
				return str
			}
		}
		str.errors[str.field] = append(str.errors[str.field], Error{ErrIn, CodeIn, values})
	}
	return str
}

func (str *Str) Date(layout string) *Str {
	if str.value != "" {
		if _, err := time.Parse(layout, str.value); err != nil {
			str.errors[str.field] = append(str.errors[str.field], Error{ErrDate, CodeDate, layout})
		}
	}
	return str
}

func (str *Str) Email() *Str {
	if str.value == "" {
		return str
	}
	return str.RegExp(PatternEmail)
}

func (str *Str) URL() *Str {
	if str.value != "" {
		return str.RegExp(PatternURL)
	}
	return str
}

func (str *Str) RegExp(pattern string) *Str {
	if str.value != "" {
		if matched, err := regexp.MatchString(pattern, str.value); err != nil {
			str.errors[str.field] = append(str.errors[str.field], Error{ErrBadParameter, CodeBadParameter, pattern})
		} else if !matched {
			str.errors[str.field] = append(str.errors[str.field], Error{ErrNotMatched, CodeNotMatched, pattern})
		}
	}
	return str
}
