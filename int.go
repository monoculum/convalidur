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

func (it *Int) Required() *Int {
	if it.value == 0 {
		it.errors[it.field] = append(it.errors[it.field], Error{ErrRequired, CodeRequired, nil})
	}
	return it
}

func (it *Int) Len(num int) *Int {
	if it.value != 0 && len(strconv.Itoa(it.value)) != num {
		it.errors[it.field] = append(it.errors[it.field], Error{ErrLen, CodeLen, num})
	}
	return it
}

func (it *Int) Equal(eq int) *Int {
	if it.value != 0 && it.value != eq {
		it.errors[it.field] = append(it.errors[it.field], Error{ErrNotEqual, CodeNotEqual, eq})
	}
	return it
}

func (it *Int) Range(min, max int) *Int {
	if it.value != 0 && (it.value < min || it.value > max) {
		it.errors[it.field] = append(it.errors[it.field], Error{ErrOutRange, CodeOutRange, []int{min, max}})
	}
	return it
}
