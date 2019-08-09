package valider

import "bytes"

type Bytes struct {
	value  []byte
	field  string
	errors Errors
}

func (v *Validator) Bytes(value []byte, field string) *Bytes {
	return &Bytes{value, field, v.Errors}
}

func (f *Bytes) Required() *Bytes {
	if len(f.value) == 0 {
		f.errors[f.field] = append(f.errors[f.field], Error{ErrRequired, CodeRequired, nil})
	}
	return f
}

func (f *Bytes) Len(num int) *Bytes {
	if len(f.value) != 0 && len(f.value) != num {
		f.errors[f.field] = append(f.errors[f.field], Error{ErrLen, CodeLen, num})
	}
	return f
}

func (f *Bytes) Range(min, max int) *Bytes {
	if len(f.value) != 0 && (len(f.value) < min || len(f.value) > max) {
		f.errors[f.field] = append(f.errors[f.field], Error{ErrOutRange, CodeOutRange, []int{min, max}})
	}
	return f
}

func (f *Bytes) In(values ...[]byte) *Bytes {
	if len(f.value) != 0 {
		for _, value := range values {
			if bytes.Equal(f.value, value) {
				return f
			}
		}
		f.errors[f.field] = append(f.errors[f.field], Error{ErrOutRange, CodeOutRange, values})
	}
	return f
}
