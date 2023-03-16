package memshare

import "testing"

func Test(t *testing.T){
	var1 := "test"

	p := Stringify[string](&var1)

	var2 := Parse[string](p)
	if var2 == nil {
		t.Error("failed to parse pointer")
		return
	}

	if *var2 != "test" {
		t.Error("pointer has the wrong value")
	}

	*var2 = "test2"

	if *var2 != "test2" {
		t.Error("var2 value failed to update")
	}

	if var1 != "test2" {
		t.Error("var1 failed to be updated by var2")
	}

	var1 = "test3"

	if var1 != "test3" {
		t.Error("var1 value failed to update")
	}

	if *var2 != "test3" {
		t.Error("var2 failed to be updated by var1")
	}
}