package valider

import "strconv"

type Int struct {
	value  int
	field  string
	errors Errors
}

func (v *Validator) Int(value int, field string) *Int {
	return &Int{value, field, v.Errors}
}

func (int *Int) Required() *Int {
	if int.value == 0 {
		int.errors[int.field] = append(int.errors[int.field], Error{ErrRequired, CodeRequired})
	}
	return int
}

func (int *Int) Len(num int) *Int {
	if int.value != 0 && len(strconv.Itoa(int.value)) != num {
		int.errors[int.field] = append(int.errors[int.field], Error{ErrLen, CodeLen})
	}
	return int
}

func (int *Int) Equal(eq int) *Int {
	if int.value != 0 && int.value != eq {
		int.errors[int.field] = append(int.errors[int.field], Error{ErrNotEqual, CodeNotEqual})
	}
	return int
}

func (int *Int) Range(min, max int) *Int {
	if int.value != 0 && (int.value < min || int.value > max) {
		int.errors[int.field] = append(int.errors[int.field], Error{ErrOutRange, CodeOutRange})
	}
	return int
}
