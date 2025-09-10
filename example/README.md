# GopherKit.Test Example Usage Guide

This directory contains comprehensive examples showing how to use all features of GopherKit.Test.

## Files Overview

- `user_service.go` - Interface definition and domain types
- `user_service_test.go` - Complete test examples using assertions, mocks, and test patterns
- `generated_example.md` - Guide for using the code generation tools

## Running the Examples

1. **Run the tests to see assertions and mocking in action:**
   ```bash
   go test ./example/ -v
   ```

2. **Generate a mock for the UserService interface:**
   ```bash
   ./gopherkit-test generate-mock ./example/user_service.go ./example/mocks/
   ```

3. **Generate test boilerplate:**
   ```bash
   ./gopherkit-test generate-test userservice ./example/generated/
   ```

4. **Generate custom assertions:**
   ```bash
   ./gopherkit-test generate-assertions ./example/generated/ \
     "IsValidEmail:email string:contains(email, \"@\") && contains(email, \".\"):invalid email format" \
     "IsPositiveInt:value int:value > 0:expected positive integer"
   ```

## Key Features Demonstrated

### 1. Comprehensive Assertions
```go
// Basic assertions
assert.Equal(t, expected, actual, "description")
assert.NotNil(t, result, "should return user")
assert.Nil(t, err, "should not return error")

// Boolean assertions
assert.True(t, condition, "condition should be true")
assert.False(t, condition, "condition should be false")
```

### 2. Advanced Mocking
```go
// Create mock
mockRepo := mock.NewMock(t)

// Set expectations
mockRepo.On("FindByID", "123").Return(expectedUser, nil)

// Verify all expectations were met
mockRepo.AssertExpectations()
```

### 3. Test Organization
- Table-driven tests for multiple scenarios
- Subtests with `t.Run()` for organized test cases
- Clear Arrange-Act-Assert structure

### 4. Error Handling Testing
- Testing both success and failure scenarios
- Validating error messages and types
- Testing edge cases and boundary conditions

## Code Generation Examples

The generated mock from `UserService` interface will include:
- Constructor function `NewUserServiceMock(t *testing.T)`
- All interface methods that delegate to the mock framework
- Convenience `OnMethodName()` functions for setting expectations
- `AssertExpectations()` method for verification

Generated test boilerplate includes:
- Proper package declarations
- Common imports (testing, assert packages)
- Basic test function templates
- Example assertion usage

Custom assertions provide:
- Domain-specific validation functions
- Consistent error messaging
- Reusable test logic
- Better test readability

## Best Practices Shown

1. **Mock Usage:**
   - Always call `AssertExpectations()` to verify mock calls
   - Use specific expectations rather than generic ones
   - Test both success and error scenarios

2. **Assertion Usage:**
   - Provide descriptive messages for failed assertions
   - Use the most specific assertion available
   - Group related assertions logically

3. **Test Organization:**
   - Use subtests for related test cases
   - Follow AAA pattern (Arrange-Act-Assert)
   - Test edge cases and error conditions
   - Use table-driven tests for multiple similar scenarios

4. **Code Generation:**
   - Generate mocks from interfaces to ensure type safety
   - Use generated test boilerplate as starting points
   - Create custom assertions for domain-specific validations
