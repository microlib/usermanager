package usermanager

// UserInterface defines an interface for any User object
type UserInterface interface {
	Id() string
	Name() string
	Email() string
	Password() string
}

// NewUser is the constructor for User
func NewUser(id string, name string, email string, password string) *User {
	return &User{id:id, name:name, email:email, password:password}
}

// User object
type User struct {
	id string
	name string
	email string
	password string
}

// Id returns the id of the User
func (u *User) Id() string {
	return u.id
}

// Name returns the name of the User
func (u *User) Name() string {
	return u.name
}

// Email returns the email of the User
func (u *User) Email() string {
	return u.email
}

// Password returns the password of the User
func (u *User) Password() string {
	return u.password
}