package bitmath

import "testing"

func TestAsMebibytes(t *testing.T) {
	testObject := Mebibyte(1)
	if testObject.AsMebibytes().Value != 1 {
		t.Fail()
	}

	testObject = Gibibyte(1)
	if testObject.AsMebibytes().Value != 1024 {
		t.Fail()
	}
}

func TestAsGibibytes(t *testing.T) {
	testObject := Gibibyte(1)
	if testObject.AsMebibytes().Value != 1024 {
		t.Fail()
	}

	testObject = Gibibyte(1)
	if testObject.AsGibibytes().Value != 1 {
		t.Fail()
	}

	testObject = Gibibyte(0.5)
	if testObject.AsMebibytes().Value != 512 {
		t.Fail()
	}

	testObject = Gibibyte(0.25)
	if testObject.AsMebibytes().Value != 256 {
		t.Fail()
	}

}

func TestAsTebibytes(t *testing.T) {
	testObject := Tebibyte(1)
	if testObject.AsGibibytes().Value != 1024 {
		t.Fail()
	}

	testObject = Gibibyte(1)
	if testObject.AsGibibytes().Value != 1 {
		t.Fail()
	}
}

func TestAsPebibytes(t *testing.T) {
	testObject := Pebibyte(1)
	if testObject.AsTebibytes().Value != 1024 {
		t.Fail()
	}

	testObject = Pebibyte(1)
	if testObject.AsPebibytes().Value != 1 {
		t.Fail()
	}
}
