package usermanager

type UsersInterface interface {
	Find(params map[string]string) ([]UserInterface, error)
	Get(id string) (*UserInterface, error)
	Create(params map[string]string)
	Update(id string, params map[string]string)
	Delete(id string)
}

type Users struct {
	users []map[string]string
}

func (u *Users) Find(params map[string]string) ([]UserInterface, error) {
	result := []*User{}
	user := NewUser("1", "Padraig", "padraig@ireland.ie", "123safe!")
	result = append(result, user)

	userInterfaces := make([]UserInterface, len(result))
    for i, d := range result {
        userInterfaces[i] = d
    }

	return userInterfaces, nil
}

func (u *Users) Get(id string) (UserInterface, error) {
    return &User{}, nil
}

func (u *Users) Create(params map[string]string) {

}

func (u *Users) Update(id string, params map[string]string) {

}

func (u *Users) Delete(id string) {

}