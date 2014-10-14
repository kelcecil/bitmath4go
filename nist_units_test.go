package bitmath

import (
	"testing"
)

func TestMebibyte(t *testing.T) {
	test := Mebibyte(1)
	if test.Prefix != "MiB" {
		t.Errorf("Prefix is %s. Should be MiB", test.Prefix)
	}
}

func TestMiB(t *testing.T) {
	shortName := MiB(1)
	longName := Mebibyte(1)
	testShortAndLongFuncNamesShouldBeSame(t, shortName, longName)
}

func TestMebibyteValuesShouldBeRight(t *testing.T) {
	testObject := Mebibyte(1)
	expectedValue := float64(1048576)
	testInitializationValues(t, testObject, expectedValue)

	testObject = Mebibyte(5)
	expectedValue = float64(5242880)
	testInitializationValues(t, testObject, expectedValue)
}

func TestGibibyte(t *testing.T) {
	test := Gibibyte(1)
	if test.Prefix != "GiB" {
		t.Errorf("Prefix is %s. Should be GiB", test.Prefix)
	}
}

func TestGib(t *testing.T) {
	shortName := GiB(1)
	longName := Gibibyte(1)
	testShortAndLongFuncNamesShouldBeSame(t, shortName, longName)
}

func TestGibibyteValuesShouldBeRight(t *testing.T) {
	testObject := Gibibyte(1)
	expectedValue := float64(1073741824)
	testInitializationValues(t, testObject, expectedValue)

	testObject = Gibibyte(5)
	expectedValue = float64(1073741824) * 5
	testInitializationValues(t, testObject, expectedValue)
}

func TestTebibyte(t *testing.T) {
	test := Tebibyte(1)
	if test.Prefix != "TiB" {
		t.Errorf("Prefix is %s. Should be TiB", test.Prefix)
	}
}

func TestTiB(t *testing.T) {
	shortName := TiB(1)
	longName := Tebibyte(1)
	testShortAndLongFuncNamesShouldBeSame(t, shortName, longName)
}

func TestTebibyteValuesShouldBeRight(t *testing.T) {
	testObject := Tebibyte(1)
	expectedValue := float64(1099511627776)
	testInitializationValues(t, testObject, expectedValue)

	testObject = Tebibyte(5)
	expectedValue = float64(1099511627776) * 5
	testInitializationValues(t, testObject, expectedValue)
}

func TestPebibyte(t *testing.T) {
	test := Pebibyte(1)
	if test.Prefix != "PiB" {
		t.Errorf("Prefix is %s. Should be PiB", test.Prefix)
	}
}

func TestPiB(t *testing.T) {
	shortName := PiB(1)
	longName := Pebibyte(1)
	testShortAndLongFuncNamesShouldBeSame(t, shortName, longName)
}

func TestPebibyteValuesShouldBeRight(t *testing.T) {
	testObject := Pebibyte(1)
	expectedValue := float64(1125899906842624)
	testInitializationValues(t, testObject, expectedValue)

	testObject = Pebibyte(5)
	expectedValue = float64(1125899906842624) * 5
	testInitializationValues(t, testObject, expectedValue)
}
