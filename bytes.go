package valider

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
		f.errors[f.field] = append(f.errors[f.field], Error{ErrRequired, CodeRequired})
	}
	return f
}

func (f *Bytes) Len(num int) *Bytes {
	if len(f.value) != 0 && len(f.value) != num {
		f.errors[f.field] = append(f.errors[f.field], Error{ErrLen, CodeLen})
	}
	return f
}

func (f *Bytes) Range(min, max int) *Bytes {
	if len(f.value) != 0 && (len(f.value) < min || len(f.value) > max) {
		f.errors[f.field] = append(f.errors[f.field], Error{ErrOutRange, CodeOutRange})
	}
	return f
}
