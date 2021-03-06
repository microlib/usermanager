package usermanager

import "errors"

type UsersInterface interface {
	Find(params map[string]string) ([]UserInterface, error)
	Get(id string) (*UserInterface, error)
	Create(params map[string]string)
	Update(id string, params map[string]string)
	Delete(id string)
}

func NewUsers(users []map[string]string) *Users {
	return &Users{users:users}
}

type Users struct {
	users []map[string]string
}

// Find returns all the users that satisfy the values passed in the map
func (u *Users) Find(params map[string]string) ([]UserInterface, error) {

	result := map[int]*User{}

	for uk, uv := range u.users { // users
		for pk, pv := range params { // params/values we're looking for
			_, isUserInResult := result[uk]
			if vv, _ := uv[pk]; vv == pv { // if user's param is the value we're looking for

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
    for _, uv := range u.users {
        if uv["id"] == id {
            return NewUser(uv["id"], uv["name"], uv["email"], uv["password"]), nil
        }
    }
    return &User{}, errors.New("User not found")
}

// Create adds a new user
func (u *Users) Create(params map[string]string) {
	u.users = append(u.users, map[string]string{
		"id":       params["id"],
		"name":     params["name"],
		"email":    params["email"],
		"password": params["password"],
	})
}

// Update modifies values of an existing user
func (u *Users) Update(id string, params map[string]string) error {
    done := false
    for _, uv := range u.users {
        if uv["id"] == id {
            done = true
            for k, v := range params {
                if _, exists := uv[k]; exists {
                    uv[k] = v
                }
            }
            // return NewUser(uv["id"], uv["name"], uv["email"], uv["password"]), nil
        }
    }
	if done == false {
        return errors.New("User #" + id + " not found")
    }
    return nil
}

// Delete removes an existing user
func (u *Users) Delete(id string) {
    for uk, uv := range u.users {
        if uv["id"] == id {
            u.users = u.removeFromArray(u.users, uk)
            break
        }
    }
}

func (u *Users) removeFromArray(s []map[string]string, i int) []map[string]string {
    s = append(s[:i], s[i+1:]...)
    return s
}
