package example

import (
	"errors"
	"testing"

	"github.com/g-restante/GopeherKit.Test/assert"
	"github.com/g-restante/GopeherKit.Test/mock"
)

// UserRepository simulates a repository for user data
type UserRepository interface {
	FindByID(id string) (*User, error)
	Save(user *User) error
	Delete(id string) error
}

// UserServiceImpl is the concrete implementation of UserService
type UserServiceImpl struct {
	repo UserRepository
}

// NewUserService creates a new service instance
func NewUserService(repo UserRepository) *UserServiceImpl {
	return &UserServiceImpl{repo: repo}
}

// GetUser implements UserService.GetUser
func (s *UserServiceImpl) GetUser(id string) (*User, error) {
	if id == "" {
		return nil, errors.New("id cannot be empty")
	}
	return s.repo.FindByID(id)
}

// CreateUser implements UserService.CreateUser
func (s *UserServiceImpl) CreateUser(name, email string) (*User, error) {
	if name == "" || email == "" {
		return nil, errors.New("name and email are required")
	}
	
	user := &User{
		ID:    "user_123", // In a real case, this would be generated
		Name:  name,
		Email: email,
	}
	
	err := s.repo.Save(user)
	if err != nil {
		return nil, err
	}
	
	return user, nil
}

// UpdateUser implements UserService.UpdateUser
func (s *UserServiceImpl) UpdateUser(id string, updates UserUpdates) error {
	if id == "" {
		return errors.New("id cannot be empty")
	}
	
	user, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	
	if updates.Name != nil {
		user.Name = *updates.Name
	}
	if updates.Email != nil {
		user.Email = *updates.Email
	}
	
	return s.repo.Save(user)
}

// DeleteUser implements UserService.DeleteUser
func (s *UserServiceImpl) DeleteUser(id string) error {
	if id == "" {
		return errors.New("id cannot be empty")
	}
	return s.repo.Delete(id)
}

// ListUsers implements UserService.ListUsers
func (s *UserServiceImpl) ListUsers(filter UserFilter, limit int) ([]*User, error) {
	// Simplified implementation - in a real case this would query the repository
	users := []*User{
		{ID: "1", Name: "John Doe", Email: "john@example.com"},
		{ID: "2", Name: "Jane Smith", Email: "jane@example.com"},
	}
	
	// Simple filtering
	var filtered []*User
	for _, user := range users {
		include := true
		
		if filter.NameContains != "" {
			if !contains(user.Name, filter.NameContains) {
				include = false
			}
		}
		
		if filter.EmailContains != "" {
			if !contains(user.Email, filter.EmailContains) {
				include = false
			}
		}
		
		if include {
			filtered = append(filtered, user)
			if len(filtered) >= limit && limit > 0 {
				break
			}
		}
	}
	
	return filtered, nil
}

// Helper function for filtering
func contains(haystack, needle string) bool {
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return true
		}
	}
	return false
}

// ---- TEST EXAMPLES ----

// TestUserService_GetUser demonstrates the use of assertions and mocking
func TestUserService_GetUser(t *testing.T) {
	// Create the mock repository
	mockRepo := mock.NewMock(t)
	service := NewUserService(&MockUserRepository{mock: mockRepo})

	t.Run("successful get user", func(t *testing.T) {
		// Arrange: Setup the mock
		expectedUser := &User{ID: "123", Name: "John Doe", Email: "john@example.com"}
		mockRepo.On("FindByID", "123").Return(expectedUser, nil)

		// Act: Call the method
		result, err := service.GetUser("123")

		// Assert: Verify the result using GopherKit.Test assertions
		assert.Nil(t, err, "Should not return error")
		assert.NotNil(t, result, "Should return user")
		assert.Equal(t, "123", result.ID, "User ID should match")
		assert.Equal(t, "John Doe", result.Name, "User name should match")
		assert.Equal(t, "john@example.com", result.Email, "User email should match")

		// Verify that all mock expectations were met
		mockRepo.AssertExpectations()
	})

	t.Run("empty id returns error", func(t *testing.T) {
		// Act
		result, err := service.GetUser("")

		// Assert: Verify that the error is handled correctly
		assert.NotNil(t, err, "Should return error for empty ID")
		assert.Nil(t, result, "Should not return user for empty ID")
		assert.Equal(t, "id cannot be empty", err.Error(), "Error message should match")
	})

	t.Run("repository error is propagated", func(t *testing.T) {
		// Arrange
		expectedError := errors.New("database connection failed")
		mockRepo.On("FindByID", "456").Return(nil, expectedError)

		// Act
		result, err := service.GetUser("456")

		// Assert
		assert.NotNil(t, err, "Should return error")
		assert.Nil(t, result, "Should not return user")
		assert.Equal(t, expectedError, err, "Should propagate repository error")

		mockRepo.AssertExpectations()
	})
}

// TestUserService_CreateUser demonstrates more complex tests
func TestUserService_CreateUser(t *testing.T) {
	mockRepo := mock.NewMock(t)
	service := NewUserService(&MockUserRepository{mock: mockRepo})

	t.Run("successful user creation", func(t *testing.T) {
		// Arrange
		mockRepo.On("Save", mock.Any).Return(nil)

		// Act
		result, err := service.CreateUser("Alice Brown", "alice@example.com")

		// Assert
		assert.Nil(t, err, "Should not return error")
		assert.NotNil(t, result, "Should return created user")
		assert.Equal(t, "user_123", result.ID, "Should have generated ID")
		assert.Equal(t, "Alice Brown", result.Name, "Name should match")
		assert.Equal(t, "alice@example.com", result.Email, "Email should match")

		mockRepo.AssertExpectations()
	})

	t.Run("validation errors", func(t *testing.T) {
		testCases := []struct {
			name     string
			userName string
			email    string
			expected string
		}{
			{"empty name", "", "test@example.com", "name and email are required"},
			{"empty email", "Test User", "", "name and email are required"},
			{"both empty", "", "", "name and email are required"},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				result, err := service.CreateUser(tc.userName, tc.email)
				
				assert.NotNil(t, err, "Should return validation error")
				assert.Nil(t, result, "Should not return user on validation error")
				assert.Equal(t, tc.expected, err.Error(), "Error message should match")
			})
		}
	})
}

// MockUserRepository is the mock for the repository (this could be generated automatically)
type MockUserRepository struct {
	mock *mock.Mock
}

func (m *MockUserRepository) FindByID(id string) (*User, error) {
	args := m.mock.Called("FindByID", id)
	var user *User
	var err error
	
	if args[0] != nil {
		user = args[0].(*User)
	}
	if args[1] != nil {
		err = args[1].(error)
	}
	
	return user, err
}

func (m *MockUserRepository) Save(user *User) error {
	args := m.mock.Called("Save", user)
	if args[0] == nil {
		return nil
	}
	return args[0].(error)
}

func (m *MockUserRepository) Delete(id string) error {
	args := m.mock.Called("Delete", id)
	if args[0] == nil {
		return nil
	}
	return args[0].(error)
}
