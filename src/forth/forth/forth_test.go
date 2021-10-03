package forth

import (
	"testing"
)

func assertEqual(t *testing.T, left, right string) {
	if left != right {
		t.Errorf("assert failed with: '%s' is not equal to '%s'", left, right)
	}
}

func TestInterpretedMath(t *testing.T) {
	f := NewForth()
	result, _ := f.Exec("1 2 + .")
	assertEqual(t, result, "3 ")

	result, _ = f.Exec("1 2 - .")
	assertEqual(t, result, "-1 ")

	result, _ = f.Exec("3 2 * .")
	assertEqual(t, result, "6 ")

	result, _ = f.Exec("6 2 / .")
	assertEqual(t, result, "3 ")
}

func TestCompiledMath(t *testing.T) {
	f := NewForth()
	result, _ := f.Exec("1 2 + .")
	assertEqual(t, result, "3 ")

	result, _ = f.Exec("1 2 - .")
	assertEqual(t, result, "-1 ")

	result, _ = f.Exec("3 2 * .")
	assertEqual(t, result, "6 ")

	result, _ = f.Exec("6 2 / .")
	assertEqual(t, result, "3 ")
}

func TestVariable(t *testing.T) {
	f := NewForth()

	result, _ := f.Exec("variable var1")
	assertEqual(t, result, "")

	result, _ = f.Exec("100 var1 !")
	assertEqual(t, result, "")

	result, _ = f.Exec("var1 @ .")
	assertEqual(t, result, "100 ")
}

func TestConstants(t *testing.T) {
	f := NewForth()
	result, _ := f.Exec("100 constant const1")
	assertEqual(t, result, "")

	result, _ = f.Exec("const1 .")
	assertEqual(t, result, "100 ")
}

func TestFunc(t *testing.T) {
	f := NewForth()
	result, _ := f.Exec(`: square dup * ;
                              3 square .`)
	assertEqual(t, result, "9 ")
}

func TestCompare(t *testing.T) {
	f := NewForth()
	result, _ := f.Exec("5 5 = .")
	assertEqual(t, result, "-1 ")

	result, _ = f.Exec("4 3 = .")
	assertEqual(t, result, "0 ")

	result, _ = f.Exec("6 5 > .")
	assertEqual(t, result, "-1 ")

	result, _ = f.Exec("6 5 < .")
	assertEqual(t, result, "0 ")

	result, _ = f.Exec("5 5 <> .")
	assertEqual(t, result, "-1 ")
}

func TestIfThen(t *testing.T) {
	f := NewForth()
	result, _ := f.Exec("1 if 2 . then")
	if result != "2 " {
		t.Fail()
	}

	result, _ = f.Exec("-1 if 3 . then")
	if result != "3 " {
		t.Fail()
	}
	result, _ = f.Exec("0 if 2 . then")
	if result != "" {
		t.Fail()
	}
}

func TestComment(t *testing.T) {
	f := NewForth()
	result, _ := f.Exec("1 ( 2 3 then loop do this is my random comment ) 4 5 .s")
	if result != "<3> 1 4 5 " {
		t.Fail()
	}
}

func TestRecurse(t *testing.T) {
	f := NewForth()
	result, _ := f.Exec(": rec 1 - dup . dup 0 > if recurse then ;")
	result, _ = f.Exec("4 rec")
	if result != "3 2 1 0 " {
		t.Fail()
	}
}

/*
func TestFaculty(t *testing.T) {
	f := NewForth()
	result := f.Exec(`
		: fac2
		dup 0> if
		dup 1- recurse *
		else
		drop 1
		endif ;
		8 fac2 .
	`)
}
*/
