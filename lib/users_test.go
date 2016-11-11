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

}
