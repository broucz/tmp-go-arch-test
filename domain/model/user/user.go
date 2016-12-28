package user

// ID uniquely identifies a particular user.
type ID string

// User is the main class of this model.
type User struct {
	ID    ID
	Name  string
	Email string
}

// GetInfo return the info of this model.
func (u *User) GetInfo() string {
	return u.Email
}

// New creates a new user.
func New(id ID, name string, email string) *User {
	return &User{
		ID:    id,
		Name:  name,
		Email: email,
	}
}

// Repository provides access to a user store.
type Repository interface {
	Store(user *User) error
	Find(id ID) (*User, error)
	FindAll() []*User
}
