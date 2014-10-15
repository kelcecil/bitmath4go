package bitmath4go

import (
	"math"
	"testing"
)

func TestKilibyte(t *testing.T) {
	test := KiB(1)
	if test.Prefix != "KiB" {
		t.Errorf("Prefix is %s. Should be KiB", test.Prefix)
	}
}

func TestKilibyteValuesShouldBeRight(t *testing.T) {
	testObject := KiB(1)
	expectedValue := math.Pow(2, 10)
	testInitializationValues(t, testObject, expectedValue)

	testObject = KiB(5)
	expectedValue = math.Pow(2, 10) * 5
	testInitializationValues(t, testObject, expectedValue)
}

func TestMebibyte(t *testing.T) {
	test := MiB(1)
	if test.Prefix != "MiB" {
		t.Errorf("Prefix is %s. Should be MiB", test.Prefix)
	}
}

func TestMebibyteValuesShouldBeRight(t *testing.T) {
	testObject := MiB(1)
	expectedValue := float64(1048576)
	testInitializationValues(t, testObject, expectedValue)

	testObject = MiB(5)
	expectedValue = float64(5242880)
	testInitializationValues(t, testObject, expectedValue)
}

func TestGibibyte(t *testing.T) {
	test := GiB(1)
	if test.Prefix != "GiB" {
		t.Errorf("Prefix is %s. Should be GiB", test.Prefix)
	}
}

func TestGibibyteValuesShouldBeRight(t *testing.T) {
	testObject := GiB(1)
	expectedValue := float64(1073741824)
	testInitializationValues(t, testObject, expectedValue)

	testObject = GiB(5)
	expectedValue = float64(1073741824) * 5
	testInitializationValues(t, testObject, expectedValue)
}

func TestTebibyte(t *testing.T) {
	test := TiB(1)
	if test.Prefix != "TiB" {
		t.Errorf("Prefix is %s. Should be TiB", test.Prefix)
	}
}

func TestTebibyteValuesShouldBeRight(t *testing.T) {
	testObject := TiB(1)
	expectedValue := float64(1099511627776)
	testInitializationValues(t, testObject, expectedValue)

	testObject = TiB(5)
	expectedValue = float64(1099511627776) * 5
	testInitializationValues(t, testObject, expectedValue)
}

func TestPebibyte(t *testing.T) {
	test := PiB(1)
	if test.Prefix != "PiB" {
		t.Errorf("Prefix is %s. Should be PiB", test.Prefix)
	}
}

func TestPebibyteValuesShouldBeRight(t *testing.T) {
	testObject := PiB(1)
	expectedValue := float64(1125899906842624)
	testInitializationValues(t, testObject, expectedValue)

	testObject = PiB(5)
	expectedValue = float64(1125899906842624) * 5
	testInitializationValues(t, testObject, expectedValue)
}

func TestExbibyte(t *testing.T) {
	test := EiB(1)
	if test.Prefix != "EiB" {
		t.Errorf("Prefix is %s. Should be EiB", test.Prefix)
	}
}

func TestExbibyteValuesShouldBeRight(t *testing.T) {
	testObject := EiB(1)
	expectedValue := math.Pow(2, 60)
	testInitializationValues(t, testObject, expectedValue)

	testObject = EiB(5)
	expectedValue = math.Pow(2, 60) * 5
	testInitializationValues(t, testObject, expectedValue)
}
