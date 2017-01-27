package usermanager

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUsers_Find(t *testing.T) {
	testCases := []struct {
		usersMap []map[string]string
		search   map[string]string
		expected []*User
	}{
		{
			[]map[string]string{
				map[string]string{
					"id":       "1",
					"name":     "Padraig",
					"email":    "padraig@ireland.ie",
					"password": "123safe!",
				},
				map[string]string{
					"id":       "2",
					"name":     "Padraig",
					"email":    "pat@superirish.com",
					"password": "mypassword123!",
				},
				map[string]string{
					"id":       "2",
					"name":     "Roberto",
					"email":    "robbie@italianovero.it",
					"password": "passwordsupersicura",
				},
			},
			map[string]string{"name": "Padraig"},
			[]*User{
				NewUser("1", "Padraig", "padraig@ireland.ie", "123safe!"),
				NewUser("2", "Padraig", "pat@superirish.com", "mypassword123!"),
			},
		},
	}
	for tk, tc := range testCases {
		t.Run(fmt.Sprintf("testing user #'%s'", tk), func(t *testing.T) {
			users := Users{tc.usersMap}
			expected := make([]UserInterface, len(tc.expected))
			for k, user := range tc.expected {
				expected[k] = user
			}

			if res, _ := users.Find(tc.search); !reflect.DeepEqual(res, expected) {
				t.Errorf("got %s; want %s", res, expected)
			}
		})
	}
}

func TestUsers_Get(t *testing.T) {
	id := "1"
	expected := NewUser("1", "Padraig", "padraig@irish.ie", "123abc")

	users := Users{[]map[string]string{
		map[string]string{
			"id":       "1",
			"name":     "Padraig",
			"email":    "padraig@irish.ie",
			"password": "123abc",
		},
		map[string]string{
			"id":       "2",
			"name":     "Padraig",
			"email":    "pat@superirish.com",
			"password": "mypassword123!",
		},
		map[string]string{
			"id":       "2",
			"name":     "Roberto",
			"email":    "robbie@italianovero.it",
			"password": "passwordsupersicura",
		},
	}}

    if result, _ := users.Get(id); !reflect.DeepEqual(result, expected) {
        t.Errorf("got %s; want %s", result, expected)
    }
}

func TestUsers_GetNotFound(t *testing.T) {
    id := "4"

    users := Users{[]map[string]string{
        map[string]string{
            "id":       "1",
            "name":     "Padraig",
            "email":    "padraig@irish.ie",
            "password": "123abc",
        },
        map[string]string{
            "id":       "2",
            "name":     "Padraig",
            "email":    "pat@superirish.com",
            "password": "mypassword123!",
        },
        map[string]string{
            "id":       "3",
            "name":     "Roberto",
            "email":    "robbie@italianovero.it",
            "password": "passwordsupersicura",
        },
    }}

    _, err := users.Get(id)

    if err == nil {
        t.Errorf("User #%s does not exist, should have returned error", id)
    }
}

func TestUsers_Delete(t *testing.T) {
    id := "3"
    users := Users{[]map[string]string{
        map[string]string{
            "id":       "1",
            "name":     "Padraig",
            "email":    "padraig@irish.ie",
            "password": "123abc",
        },
        map[string]string{
            "id":       "2",
            "name":     "Padraig",
            "email":    "pat@superirish.com",
            "password": "mypassword123!",
        },
        map[string]string{
            "id":       "3",
            "name":     "Roberto",
            "email":    "robbie@italianovero.it",
            "password": "passwordsupersicura",
        },
    }}

    users.Delete(id)

    if _, err := users.Get(id); err == nil {
        t.Errorf("User #%s should have been deleted", id)
    }
}

func TestUsers_Update(t *testing.T) {
    id := "3"
    users := Users{[]map[string]string{
        map[string]string{
            "id":       "1",
            "name":     "Padraig",
            "email":    "padraig@irish.ie",
            "password": "123abc",
        },
        map[string]string{
            "id":       "2",
            "name":     "Padraig",
            "email":    "pat@superirish.com",
            "password": "mypassword123!",
        },
        map[string]string{
            "id":       "3",
            "name":     "Roberto",
            "email":    "robbie@italianovero.it",
            "password": "passwordsupersicura",
        },
    }}
    params := map[string]string{
        "name": "Carlo",
        "email": "carlo.bonin@tantaroba.com",
    }

    users.Update(id, params)
    user, _ := users.Get(id)

    if (user.Name() != "Carlo" || user.Email() != "carlo.bonin@tantaroba.com") {
        t.Errorf("User #%s should have been modified", id)
    }
}
