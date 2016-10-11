package usermanager

type UserInterface interface {
	Id() string
	Name() string
	Email() string
	Password() string
}

func NewUser(id string, name string, email string, password string) *User {
	return &User{id:id, name:name, email:email, password:password}
}

type User struct {
	id string
	name string
	email string
	password string
}

func (u *User) Id() string {
	return u.id
}

func (u *User) Name() string {
	return u.name
}

func (u *User) Email() string {
	return u.email
}

func (u *User) Password() string {
	return u.password
}