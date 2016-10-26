package usermanager

type UsersInterface interface {
	Find(params map[string]string) []*UserInterface
	Get(id string) *UserInterface
	Create(params map[string]string)
	Update(id string, params map[string]string)
	Delete(id string)
}

type Users struct {
	users []map[string]string
}

func (u *Users) Find(params map[string]string) []*UserInterface {

}
