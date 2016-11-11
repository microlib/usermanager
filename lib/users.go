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

// Find returns all the users that satisfy the values passed in the map
func (u *Users) Find(params map[string]string) ([]UserInterface, error) {

    result := map[int]*User{}

    for uk, uv := range u.users {               // users
        for pk, pv := range params {            // params/values we're looking for
            _, isUserInResult := result[uk]
            if vv, _ := uv[pk]; vv == pv {      // if user's param is the value we're looking for

                // if user is not in "result", add it
				if !isUserInResult {
                    result[uk] = NewUser(uv["id"], uv["name"], uv["email"], uv["password"])
                }
                break
            } else {
                // here if user is in "result", remove it
                if isUserInResult {
                    delete(result, uk)
                }
            }


        }
    }

	userInterfaces := make([]UserInterface, len(result))
	for i, d := range result {
		userInterfaces[i] = d
	}

	return userInterfaces, nil
}


// Get returns a single user by ID
func (u *Users) Get(id string) (UserInterface, error) {
	return NewUser("1", "Padraig", "padraig@irish.ie", "123abc"), nil
}

// Create adds a new user
func (u *Users) Create(params map[string]string) {

}

// Update modifies values of an existing user
func (u *Users) Update(id string, params map[string]string) {

}

// Delete removes an existing user
func (u *Users) Delete(id string) {

}
