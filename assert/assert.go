package assert

import (
	"reflect"
	"testing"
)

// Equal asserts that two values are equal. If they are not equal, it calls t.Errorf.
// The optional msg parameter allows for a custom error message.
func Equal(t *testing.T, expected, actual any, msg ...string) {
	t.Helper()
	
	if !reflect.DeepEqual(expected, actual) {
		var message string
		if len(msg) > 0 && msg[0] != "" {
			message = msg[0]
		} else {
			message = "values should be equal"
		}
		
		t.Errorf("%s\nExpected: %v\nActual:   %v", message, expected, actual)
	}
}

// NotEqual asserts that two values are not equal. If they are equal, it calls t.Errorf.
func NotEqual(t *testing.T, expected, actual any, msg ...string) {
	t.Helper()
	
	if reflect.DeepEqual(expected, actual) {
		var message string
		if len(msg) > 0 && msg[0] != "" {
			message = msg[0]
		} else {
			message = "values should not be equal"
		}
		
		t.Errorf("%s\nBoth values: %v", message, expected)
	}
}

// True asserts that the given value is true.
func True(t *testing.T, value bool, msg ...string) {
	t.Helper()
	
	if !value {
		var message string
		if len(msg) > 0 && msg[0] != "" {
			message = msg[0]
		} else {
			message = "expected true but got false"
		}
		
		t.Errorf(message)
	}
}

// False asserts that the given value is false.
func False(t *testing.T, value bool, msg ...string) {
	t.Helper()
	
	if value {
		var message string
		if len(msg) > 0 && msg[0] != "" {
			message = msg[0]
		} else {
			message = "expected false but got true"
		}
		
		t.Errorf(message)
	}
}

// Nil asserts that the given value is nil.
func Nil(t *testing.T, value any, msg ...string) {
	t.Helper()
	
	if value != nil && !reflect.ValueOf(value).IsNil() {
		var message string
		if len(msg) > 0 && msg[0] != "" {
			message = msg[0]
		} else {
			message = "expected nil value"
		}
		
		t.Errorf("%s\nGot: %v", message, value)
	}
}

// NotNil asserts that the given value is not nil.
func NotNil(t *testing.T, value any, msg ...string) {
	t.Helper()
	
	if value == nil || (reflect.ValueOf(value).Kind() == reflect.Ptr && reflect.ValueOf(value).IsNil()) {
		var message string
		if len(msg) > 0 && msg[0] != "" {
			message = msg[0]
		} else {
			message = "expected non-nil value"
		}
		
		t.Errorf(message)
	}
}