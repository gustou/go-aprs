package aprs

import (
	"fmt"
	"testing"
)

func TestAddressFromString(t *testing.T) {
	testCases := []struct {
		addressString string
		expected      Address
	}{
		{addressString: "K6LRG-C", expected: Address{Call: "K6LRG", SSID: "C"}},
		{addressString: "K6LRG", expected: Address{Call: "K6LRG", SSID: ""}},
		{addressString: "", expected: Address{Call: "", SSID: ""}},
		{addressString: "K6LRG-C-C", expected: Address{Call: "K6LRG", SSID: "C-C"}},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			if got := AddressFromString(testCase.addressString); got != testCase.expected {
				t.Fatalf("AddressFromString(%q) got %v, expected %v", testCase.addressString, got, testCase.expected)
			}
		})
	}
}

func TestCallPass(t *testing.T) {
	testCases := []struct {
		call     Address
		expected int16
	}{
		{call: Address{Call: "KG6HWF", SSID: "9"}, expected: 22955},
		{call: Address{Call: "KG6HWF", SSID: ""}, expected: 22955},
		{call: Address{Call: "KE6AFE", SSID: "13"}, expected: 18595},
		{call: Address{Call: "KE6AFE", SSID: ""}, expected: 18595},
		{call: Address{Call: "K6MGD", SSID: "10"}, expected: 12691},
		{call: Address{Call: "K6MGD", SSID: ""}, expected: 12691},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			if got := testCase.call.CallPass(); got != testCase.expected {
				t.Fatalf("%v.CallPass() got %d, expected %d", testCase.call, got, testCase.expected)
			}
		})
	}
}
