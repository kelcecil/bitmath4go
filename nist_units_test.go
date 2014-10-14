package bitmath4go

import (
	"math"
	"testing"
)

func TestKilibyte(t *testing.T) {
	test := Kilibyte(1)
	if test.Prefix != "KiB" {
		t.Errorf("Prefix is %s. Should be KiB", test.Prefix)
	}
}

func TestKiB(t *testing.T) {
	shortName := KiB(1)
	longName := Kilibyte(1)
	testShortAndLongFuncNamesShouldBeSame(t, shortName, longName)
}

func TestKilibyteValuesShouldBeRight(t *testing.T) {
	testObject := Kilibyte(1)
	expectedValue := math.Pow(2, 10)
	testInitializationValues(t, testObject, expectedValue)

	testObject = Kilibyte(5)
	expectedValue = math.Pow(2, 10) * 5
	testInitializationValues(t, testObject, expectedValue)
}

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

func TestExbibyte(t *testing.T) {
	test := Kilibyte(1)
	if test.Prefix != "KiB" {
		t.Errorf("Prefix is %s. Should be KiB", test.Prefix)
	}
}

func TestEiB(t *testing.T) {
	shortName := EiB(1)
	longName := Exbibyte(1)
	testShortAndLongFuncNamesShouldBeSame(t, shortName, longName)
}

func TestExbibyteValuesShouldBeRight(t *testing.T) {
	testObject := Exbibyte(1)
	expectedValue := math.Pow(2, 60)
	testInitializationValues(t, testObject, expectedValue)

	testObject = Exbibyte(5)
	expectedValue = math.Pow(2, 60) * 5
	testInitializationValues(t, testObject, expectedValue)
}
