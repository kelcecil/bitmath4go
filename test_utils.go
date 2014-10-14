package bitmath

import "testing"

func testShortAndLongFuncNamesShouldBeSame(t *testing.T, shortName Byte, longName Byte) {
	if shortName != longName {
		t.Errorf("%s short and long func names do not produce the same thing.", shortName.Prefix)
	}
}

func testInitializationValues(t *testing.T, testObject Byte, expectedValue float64) {
	if testObject.ByteValue != expectedValue {
		reportIncorrectValue(t, testObject.Prefix, testObject.ByteValue, expectedValue)
	}
}

func reportIncorrectValue(t *testing.T, size string, actual float64, expected float64) {
	t.Errorf("%s value is incorrect. Actual: %f", size, actual)
}
