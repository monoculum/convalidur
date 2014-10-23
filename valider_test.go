package valider

import "testing"

func TestStringRequired(t *testing.T) {
	str := ""
	errors := make(map[string][]string)
	valid := New(&errors)
	valid.Str(str, "Required").Required()
	if len(errors) == 0 {
		t.Errorf("should be empty")
	}

	str = "hola"
	errors = make(map[string][]string)
	valid = New(&errors)
	valid.Str(str, "Required").Required()
	if len(errors) > 0 {
		t.Errorf("required")
	}
}

func TestStringEqual(t *testing.T) {
	str := "hola"
	errors := make(map[string][]string)
	valid := New(&errors)
	valid.Str(str, "Required").Equal("adios")
	if len(errors) == 0 {
		t.Errorf("there should be errors...")
	}

	str = "hola"
	errors = make(map[string][]string)
	valid = New(&errors)
	valid.Str(str, "Required").Equal("hola")
	if len(errors) > 0 {
		t.Errorf("there should not be errors...")
	}
}

func TestStringLen(t *testing.T) {
	str := "hola"
	errors := make(map[string][]string)
	valid := New(&errors)
	valid.Str(str, "Required").Len(5)
	if len(errors) == 0 {
		t.Errorf("there should be errors...")
	}

	str = "hola"
	errors = make(map[string][]string)
	valid = New(&errors)
	valid.Str(str, "Required").Len(4)
	if len(errors) > 0 {
		t.Errorf("there should not be errors...")
	}
}
