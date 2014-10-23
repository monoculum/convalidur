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
	valid.Str(str, "Equal").Equal("adios")
	if len(errors) == 0 {
		t.Errorf("there should be errors...")
	}

	str = "hola"
	errors = make(map[string][]string)
	valid = New(&errors)
	valid.Str(str, "Equal").Equal("hola")
	if len(errors) > 0 {
		t.Errorf("there should not be errors...")
	}
}

func TestStringLen(t *testing.T) {
	str := "hola"
	errors := make(map[string][]string)
	valid := New(&errors)
	valid.Str(str, "Len").Len(5)
	if len(errors) == 0 {
		t.Errorf("there should be errors...")
	}

	str = "hola"
	errors = make(map[string][]string)
	valid = New(&errors)
	valid.Str(str, "Len").Len(4)
	if len(errors) > 0 {
		t.Errorf("there should not be errors...")
	}
}

func TestSliceIn(t *testing.T) {
	slice := []string{"hola", "adios"}
	errors := make(map[string][]string)
	valid := New(&errors)
	valid.Slice(slice, "In").In([]string{"hi", "bye"})
	if len(errors) == 0 {
		t.Errorf("there should be errors...")
	}

	slice = []string{"hola", "adios"}
	errors = make(map[string][]string)
	valid = New(&errors)
	valid.Slice(slice, "In").In([]string{"hola", "adios"})
	if len(errors) > 0 {
		t.Errorf("there should not be errors...")
	}
}

func TestMapRequired(t *testing.T) {
	ma := map[string]string{}
	errors := make(map[string][]string)
	valid := New(&errors)
	valid.Map(ma, "Required").Required()
	if len(errors) == 0 {
		t.Errorf("there should be errors...")
	}

	ma = map[string]string{"hola": "adios"}
	errors = make(map[string][]string)
	valid = New(&errors)
	valid.Map(ma, "Required").Required()
	if len(errors) > 0 {
		t.Errorf("there should not be errors...")
	}
}

func TestMapRange(t *testing.T) {
	ma := map[string]string{"hola": "adios"}
	errors := make(map[string][]string)
	valid := New(&errors)
	valid.Map(ma, "Required").Range(3, 4)
	if len(errors) == 0 {
		t.Errorf("there should be errors...")
	}

	ma = map[string]string{"hola": "adios", "hi": "bye", "salut": "bye bye"}
	errors = make(map[string][]string)
	valid = New(&errors)
	valid.Map(ma, "Required").Range(3, 4)
	if len(errors) > 0 {
		t.Errorf("there should not be errors...")
	}
}
