package mock

import (
	"fmt"
	"reflect"
	"testing"
)

// Any is a placeholder that matches any argument in mock expectations.
var Any = &anyMatcher{}

type anyMatcher struct{}

func (a *anyMatcher) String() string {
	return "mock.Any"
}

// Mock represents a mock object for testing.
type Mock struct {
	t         *testing.T
	calls     []*Call
	callCount map[string]int
}

// Call represents a mocked method call with its expected arguments and return values.
type Call struct {
	methodName string
	args       []any
	returns    []any
	called     bool
	callCount  int
}

// NewMock creates a new mock object.
func NewMock(t *testing.T) *Mock {
	return &Mock{
		t:         t,
		calls:     make([]*Call, 0),
		callCount: make(map[string]int),
	}
}

// On sets up an expectation for a method call with the given arguments.
func (m *Mock) On(methodName string, args ...any) *Call {
	call := &Call{
		methodName: methodName,
		args:       args,
		returns:    make([]any, 0),
	}
	m.calls = append(m.calls, call)
	return call
}

// Return sets the return values for the mocked method call.
func (c *Call) Return(values ...any) *Call {
	c.returns = values
	return c
}

// Times sets the expected number of times this method should be called.
func (c *Call) Times(count int) *Call {
	// Implementation for call count verification would go here
	return c
}

// Once is a convenience method that sets the expected call count to 1.
func (c *Call) Once() *Call {
	return c.Times(1)
}

// Called marks this call as having been invoked and returns the configured return values.
func (m *Mock) Called(methodName string, args ...any) []any {
	m.t.Helper()
	
	// Find matching call
	for _, call := range m.calls {
		if call.methodName == methodName && m.argsMatch(call.args, args) {
			call.called = true
			call.callCount++
			m.callCount[methodName]++
			return call.returns
		}
	}
	
	// No matching call found
	m.t.Errorf("Unexpected call to %s with args: %v", methodName, args)
	return nil
}

// AssertExpectations verifies that all expected method calls were made.
func (m *Mock) AssertExpectations() {
	m.t.Helper()
	
	for _, call := range m.calls {
		if !call.called {
			m.t.Errorf("Expected call to %s with args %v was not made", call.methodName, call.args)
		}
	}
}

// argsMatch compares two slices of arguments for equality.
func (m *Mock) argsMatch(expected, actual []any) bool {
	if len(expected) != len(actual) {
		return false
	}
	
	for i, expectedArg := range expected {
		// Check if expected argument is mock.Any
		if _, isAny := expectedArg.(*anyMatcher); isAny {
			continue // mock.Any matches any value
		}
		
		if !reflect.DeepEqual(expectedArg, actual[i]) {
			return false
		}
	}
	
	return true
}

// Reset clears all call expectations and history.
func (m *Mock) Reset() {
	m.calls = make([]*Call, 0)
	m.callCount = make(map[string]int)
}

// GetCallCount returns the number of times a method was called.
func (m *Mock) GetCallCount(methodName string) int {
	return m.callCount[methodName]
}

// String returns a string representation of the mock for debugging.
func (m *Mock) String() string {
	return fmt.Sprintf("Mock with %d expected calls", len(m.calls))
}
