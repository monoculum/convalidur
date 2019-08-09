package valider

type Bool struct {
	value  bool
	field  string
	errors Errors
}

func (v *Validator) Bool(value bool, field string) *Bool {
	return &Bool{value, field, v.Errors}
}

func (b *Bool) Equal(v bool) *Bool {
	if b.value != v {
		b.errors[b.field] = append(b.errors[b.field], Error{ErrNotEqual, CodeNotEqual, v})
	}
	return b
}

func (b *Bool) Valid() bool {
	return len(b.errors) == 0
}
