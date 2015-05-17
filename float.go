package valider

import "strconv"

type Float struct {
	value  float64
	field  string
	errors *Errors
}

func (v *Validator) Float(value float64, field string) *Float {
	return &Float{value, field, v.Errors}
}

func (f *Float) Required() *Float {
	if f.value == 0 {
		(*f.errors)[f.field] = append((*f.errors)[f.field], Error{ErrRequired, CodeRequired})
	}
	return f
}

func (f *Float) Len(num int) *Float {
	if f.value != 0 && len(strconv.FormatFloat(f.value, 'f', -1, 64)) != num {
		(*f.errors)[f.field] = append((*f.errors)[f.field], Error{ErrLen, CodeLen})
	}
	return f
}

func (f *Float) Equal(eq float64) *Float {
	if f.value != 0 && f.value != eq {
		(*f.errors)[f.field] = append((*f.errors)[f.field], Error{ErrNotEqual, CodeNotEqual})
	}
	return f
}

func (f *Float) Range(min, max float64) *Float {
	if f.value != 0 && (f.value < min || f.value > max) {
		(*f.errors)[f.field] = append((*f.errors)[f.field], Error{ErrOutRange, CodeOutRange})
	}
	return f
}
