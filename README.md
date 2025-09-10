## GopherKit.Test 
GopherKit.Test is a comprehensive testing framework designed to enhance your Go testing experience. It provides comprehensive assertion utilities, flexible mocking capabilities, and code generation tools to streamline your testing workflow.

Built with Go's native testing package as its foundation, GopherKit.Test extends the standard library with modern testing patterns and automation tools that help developers write better, more maintainable tests.

## Features

- **Comprehensive Assertions**: Rich set of assertion functions for common testing scenarios including equality checks, nil/non-nil validation, boolean assertions, and more
- **Flexible Mocking**: Easy-to-use mocking system for interfaces and dependencies with expectation verification and flexible parameter matching
- **Code Generation**: Automatic generation of test boilerplate, mock implementations, and custom assertions through powerful AST parsing
- **Clean API**: Intuitive and readable testing syntax that integrates seamlessly with Go's testing package
- **Command-line Tools**: Built-in CLI for code generation tasks that speeds up development workflow
- **Type Safety**: Full type safety with compile-time checks and runtime validation
- **Zero Dependencies**: Built entirely on Go's standard library for maximum compatibility

## Installation

Install GopherKit.Test using Go modules:

```bash
go get github.com/g-restante/GopeherKit.Test
```

### Requirements

- Go 1.21 or later
- Standard Go toolchain

## Quick Start

### Basic Assertions

GopherKit.Test provides a comprehensive set of assertion functions that make your tests more readable and maintainable:

```go
package main

import (
    "testing"
    "github.com/g-restante/GopeherKit.Test/assert"
)

func TestAssertions(t *testing.T) {
    // Basic equality
    assert.Equal(t, 4, 2+2, "Addition should work correctly")
    assert.NotEqual(t, 5, 2+2, "Should not be equal")
    
    // Boolean assertions
    assert.True(t, true, "Should be true")
    assert.False(t, false, "Should be false")
    
    // Nil checks
    var ptr *string
    assert.Nil(t, ptr, "Pointer should be nil")
    
    str := "hello"
    assert.NotNil(t, &str, "Pointer should not be nil")
}
```

### Advanced Mocking

Create flexible mocks with expectation verification and parameter matching:

```go
package main

import (
    "testing"
    "github.com/g-restante/GopeherKit.Test/mock"
)

// Example interface to mock
type UserRepository interface {
    GetUser(id int) (*User, error)
    CreateUser(name string, email string) (*User, error)
}

type User struct {
    ID    int
    Name  string
    Email string
}

func TestUserService(t *testing.T) {
    // Create a new mock
    mockRepo := mock.NewMock(t)
    
    // Set up expectations with specific parameters
    expectedUser := &User{ID: 1, Name: "John", Email: "john@example.com"}
    mockRepo.On("GetUser", 1).Return(expectedUser, nil)
    
    // Set up expectations with flexible parameter matching
    mockRepo.On("CreateUser", mock.Any, mock.Any).Return(&User{ID: 2}, nil)
    
    // Use your mock in the service
    // service := NewUserService(mockRepo)
    // user, err := service.GetUserByID(1)
    
    // Verify all expectations were met
    mockRepo.AssertExpectations(t)
}
```

### Code Generation

GopherKit.Test includes powerful code generation capabilities that analyze your Go code using AST parsing to automatically create boilerplate code, reducing repetitive tasks and ensuring consistency across your test suite.

#### Building the CLI Tool

First, build the command-line tool:

```bash
# Build the CLI tool
go build -o gopherkit-test ./cmd/gopherkittest/

# Verify installation
./gopherkit-test --help
```

#### Generate Mock from Interface

Automatically generate mock implementations from Go interfaces:

```bash
# Basic usage: generate mock for an interface
./gopherkit-test generate-mock <interface_file> <output_directory>

# Example: Generate mock for UserService interface
./gopherkit-test generate-mock ./example/user_service.go ./mocks/

# This creates a MockUserRepository struct with all interface methods
# The generated mock includes:
# - Method implementations with call tracking
# - Expectation setting with On() method
# - Return value configuration with Return() method
# - Automatic verification of call expectations
```

Generated mock example:
```go
type MockUserRepository struct {
    mock.Mock
}

func (m *MockUserRepository) GetUser(id int) (*User, error) {
    args := m.Called(id)
    return args.Get(0).(*User), args.Error(1)
}
```

#### Generate Test Boilerplate

Create structured test files with common patterns:

```bash
# Generate test file template
./gopherkit-test generate-test <package_name> <output_directory>

# Example
./gopherkit-test generate-test mypackage ./tests/
```

This generates:
- Table-driven test structure
- Subtest organization
- Setup and teardown patterns
- Import statements for assertions and mocks

#### Generate Custom Assertions

Create domain-specific assertion functions tailored to your needs:

```bash
# Generate custom assertion functions
./gopherkit-test generate-assertions <output_directory> "<assertion_spec>"

# Examples:
./gopherkit-test generate-assertions ./assert/ "IsPositive:value int:value > 0:expected positive value"
./gopherkit-test generate-assertions ./assert/ "IsEmpty:s string:len(s) == 0:expected empty string"
./gopherkit-test generate-assertions ./assert/ "Contains:slice []string, item string:containsString(slice, item):slice should contain item"
```

**Assertion Specification Format:**
`"name:parameters:condition:defaultMessage"`

- **name**: Function name (e.g., `IsPositive`)
- **parameters**: Function parameters with types (e.g., `value int`)
- **condition**: Boolean condition to check (e.g., `value > 0`)
- **defaultMessage**: Default error message (e.g., `expected positive value`)

Generated assertion example:
```go
func IsPositive(t *testing.T, value int, msgAndArgs ...interface{}) {
    t.Helper()
    if !(value > 0) {
        message := "expected positive value"
        if len(msgAndArgs) > 0 {
            message = fmt.Sprintf(msgAndArgs[0].(string), msgAndArgs[1:]...)
        }
        t.Errorf("IsPositive assertion failed: %s", message)
    }
}
```
```

## API Reference

### Assertions (`github.com/g-restante/GopeherKit.Test/assert`)

| Function | Description | Example |
|----------|-------------|---------|
| `Equal(t, expected, actual, msgAndArgs...)` | Asserts that two values are equal | `assert.Equal(t, 42, result)` |
| `NotEqual(t, expected, actual, msgAndArgs...)` | Asserts that two values are not equal | `assert.NotEqual(t, 0, len(slice))` |
| `True(t, value, msgAndArgs...)` | Asserts that a value is true | `assert.True(t, isValid)` |
| `False(t, value, msgAndArgs...)` | Asserts that a value is false | `assert.False(t, hasError)` |
| `Nil(t, value, msgAndArgs...)` | Asserts that a value is nil | `assert.Nil(t, err)` |
| `NotNil(t, value, msgAndArgs...)` | Asserts that a value is not nil | `assert.NotNil(t, user)` |

### Mocking (`github.com/g-restante/GopeherKit.Test/mock`)

#### Mock Methods

| Method | Description | Example |
|--------|-------------|---------|
| `NewMock(t)` | Creates a new mock instance | `m := mock.NewMock(t)` |
| `On(methodName, args...)` | Sets up method expectation | `m.On("GetUser", 123)` |
| `Return(values...)` | Sets return values for expectation | `m.On("GetUser", 123).Return(user, nil)` |
| `Called(args...)` | Records method call and returns configured values | `return m.Called(id)` |
| `AssertExpectations(t)` | Verifies all expectations were met | `m.AssertExpectations(t)` |

#### Special Matchers

| Matcher | Description | Example |
|---------|-------------|---------|
| `mock.Any` | Matches any value of any type | `m.On("Method", mock.Any)` |

### Code Generation (`./gopherkit-test`)

#### Commands

| Command | Description | Syntax |
|---------|-------------|---------|
| `generate-mock` | Generate mock from interface | `./gopherkit-test generate-mock <file> <output>` |
| `generate-test` | Generate test boilerplate | `./gopherkit-test generate-test <package> <output>` |
| `generate-assertions` | Generate custom assertions | `./gopherkit-test generate-assertions <output> <spec>` |

## Examples

### Complete Test Suite Example

Here's a comprehensive example showing how to use all features together:

```go
package example

import (
    "testing"
    "github.com/g-restante/GopeherKit.Test/assert"
    "github.com/g-restante/GopeherKit.Test/mock"
)

// Service under test
type UserService struct {
    repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
    return &UserService{repo: repo}
}

func (s *UserService) CreateUser(name, email string) (*User, error) {
    if name == "" {
        return nil, errors.New("name cannot be empty")
    }
    return s.repo.CreateUser(name, email)
}

// Test suite with table-driven tests
func TestUserService(t *testing.T) {
    tests := []struct {
        name        string
        userName    string
        userEmail   string
        mockSetup   func(*mock.Mock)
        expectError bool
        expectUser  *User
    }{
        {
            name:      "successful user creation",
            userName:  "John Doe",
            userEmail: "john@example.com",
            mockSetup: func(m *mock.Mock) {
                m.On("CreateUser", "John Doe", "john@example.com").
                  Return(&User{ID: 1, Name: "John Doe", Email: "john@example.com"}, nil)
            },
            expectError: false,
            expectUser:  &User{ID: 1, Name: "John Doe", Email: "john@example.com"},
        },
        {
            name:        "empty name validation",
            userName:    "",
            userEmail:   "john@example.com",
            mockSetup:   func(m *mock.Mock) {}, // No mock expectations for validation failure
            expectError: true,
            expectUser:  nil,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Setup
            mockRepo := mock.NewMock(t)
            tt.mockSetup(mockRepo)
            service := NewUserService(mockRepo)

            // Execute
            user, err := service.CreateUser(tt.userName, tt.userEmail)

            // Assert
            if tt.expectError {
                assert.NotNil(t, err, "Expected an error")
                assert.Nil(t, user, "User should be nil when error occurs")
            } else {
                assert.Nil(t, err, "Should not return an error")
                assert.NotNil(t, user, "User should not be nil")
                assert.Equal(t, tt.expectUser.Name, user.Name, "User name should match")
                assert.Equal(t, tt.expectUser.Email, user.Email, "User email should match")
            }

            // Verify mock expectations
            mockRepo.AssertExpectations(t)
        })
    }
}
```

### Integration with Go's Testing Tools

GopherKit.Test works seamlessly with standard Go testing tools:

```bash
# Run tests with coverage
go test -cover ./...

# Run tests with verbose output
go test -v ./...

# Run specific test
go test -run TestUserService ./...

# Benchmark tests (if you have benchmarks)
go test -bench=. ./...
```

## Best Practices

### 1. Test Organization

```go
func TestServiceName(t *testing.T) {
    t.Run("MethodName", func(t *testing.T) {
        t.Run("successful_case", func(t *testing.T) {
            // Test implementation
        })
        
        t.Run("error_case", func(t *testing.T) {
            // Test implementation
        })
    })
}
```

### 2. Mock Usage

- Use `mock.Any` sparingly - prefer specific parameter matching
- Always call `AssertExpectations(t)` at the end of tests
- Set up mocks in the order methods will be called
- Use descriptive mock variable names (e.g., `mockUserRepo`, `mockEmailService`)

### 3. Assertion Messages

```go
// Good: descriptive messages
assert.Equal(t, expectedUser.ID, actualUser.ID, "User ID should match after creation")

// Avoid: generic or no messages
assert.Equal(t, expectedUser.ID, actualUser.ID)
```

## Architecture

### Package Structure

```
GopherKit.Test/
├── assert/           # Assertion functions
│   └── assert.go
├── mock/            # Mocking framework
│   └── mock.go
├── internal/        # Code generation engine
│   ├── generator.go
│   └── generator_test.go
├── cmd/            # CLI tool
│   └── gopherkittest/
│       └── main.go
├── example/        # Usage examples
│   ├── user_service.go
│   ├── user_service_impl.go
│   └── user_service_test.go
└── README.md
```

### Design Principles

1. **Minimalism**: Built on Go's standard library with zero external dependencies
2. **Compatibility**: Full compatibility with Go's native `testing` package
3. **Type Safety**: Compile-time and runtime type checking
4. **Extensibility**: Easy to extend with custom assertions and matchers
5. **Performance**: Minimal overhead and efficient execution

## Troubleshooting

### Common Issues

#### Mock Expectations Not Met
```bash
Error: mock: I don't know what to return because the method call was unexpected.
```
**Solution**: Ensure all method calls on mocks have corresponding `On()` expectations set up.

#### Assertion Failures
```bash
Error: Equal assertion failed: expected 42, got 24
```
**Solution**: Check that your expected and actual values match. Use debug prints if necessary.

#### Code Generation Errors
```bash
Error: failed to parse Go file: expected 'package', found 'EOF'
```
**Solution**: Ensure the target file is valid Go code and the path is correct.

### FAQ

**Q: Can I use GopherKit.Test with other testing frameworks?**
A: Yes, GopherKit.Test is designed to be compatible with any testing framework that uses Go's standard `*testing.T` type.

**Q: How do I create custom matchers for mocks?**
A: Currently, `mock.Any` is the only built-in matcher. You can extend the framework by implementing custom matcher types.

**Q: Can I generate mocks for structs, not just interfaces?**
A: The current version only supports interface mocking, which follows Go's best practices for testable code design.

**Q: Is GopherKit.Test thread-safe?**
A: Mock objects are not thread-safe. Each test should use its own mock instances.

## Performance Considerations

- Assertion functions use `t.Helper()` to ensure proper stack traces
- Mock call tracking has minimal overhead
- Code generation happens at development time, not runtime
- All reflection is done during mock setup, not during test execution

## Roadmap

### Version 1.x (Current)
- ✅ Core assertion library
- ✅ Basic mocking framework
- ✅ Code generation tools
- ✅ CLI interface

### Version 2.x (Planned)
- [ ] Advanced matchers (regex, type-based, custom)
- [ ] HTTP testing utilities
- [ ] Database testing helpers
- [ ] Performance benchmarking tools
- [ ] IDE integration plugins

### Version 3.x (Future)
- [ ] Fuzzing integration
- [ ] Property-based testing
- [ ] Visual test reporting
- [ ] CI/CD integration tools

## Contributing

We welcome contributions from the community! Here's how you can help:

### Ways to Contribute

1. **Bug Reports**: Open an issue with detailed reproduction steps
2. **Feature Requests**: Suggest new functionality with use cases
3. **Code Contributions**: Submit pull requests for bug fixes or features
4. **Documentation**: Improve README, examples, or code comments
5. **Testing**: Help test new features and report issues

### Development Setup

```bash
# Clone the repository
git clone https://github.com/g-restante/GopeherKit.Test.git
cd GopeherKit.Test

# Run tests
go test ./...

# Build CLI tool
go build -o gopherkit-test ./cmd/gopherkittest/

# Run examples
go test ./example/ -v
```

### Contribution Guidelines

1. **Code Style**: Follow standard Go formatting (`gofmt`)
2. **Testing**: Add tests for new functionality
3. **Documentation**: Update README and code comments
4. **Commits**: Write clear, descriptive commit messages
5. **Pull Requests**: Include description of changes and test results

### Code of Conduct

- Be respectful and inclusive
- Provide constructive feedback
- Focus on the technical aspects of contributions
- Help create a welcoming environment for all contributors

## License

This project is licensed under the terms found in the LICENSE file.

## Support

- **Issues**: Report bugs and request features on GitHub Issues
- **Discussions**: Join community discussions on GitHub Discussions  
- **Email**: For private inquiries, contact the maintainers

---

Made with ❤️ for the Go community. Happy testing!