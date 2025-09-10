package example

// UserService represents a service for managing users.
type UserService interface {
	// GetUser retrieves a user by ID.
	GetUser(id string) (*User, error)
	
	// CreateUser creates a new user.
	CreateUser(name, email string) (*User, error)
	
	// UpdateUser updates an existing user.
	UpdateUser(id string, updates UserUpdates) error
	
	// DeleteUser deletes a user by ID.
	DeleteUser(id string) error
	
	// ListUsers lists all users with optional filtering.
	ListUsers(filter UserFilter, limit int) ([]*User, error)
}

// User represents a user in the system.
type User struct {
	ID    string
	Name  string
	Email string
}

// UserUpdates represents the fields that can be updated for a user.
type UserUpdates struct {
	Name  *string
	Email *string
}

// UserFilter represents filtering options for listing users.
type UserFilter struct {
	NameContains  string
	EmailContains string
}
