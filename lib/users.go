package usermanager

type UsersInterface interface {
	Find() []*UserInterface
	FindOne() *UserInterface
	Create()
	Update()
	Delete()
}

type Users struct {
}

