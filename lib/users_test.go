package usermanager

import (
    "testing"
)

func TestUsers_Find(t *testing.T) {
    //testCases := []struct {
    //    params map[string]string
    //    key    string
    //    value  string
    //}{
    //    {map[string]string{"hello": "test"}, "hell", "Value not found"},
    //    {map[string]string{"hello": "test"}, "hi", "Value not found"},
    //}
    //for _, tc := range testCases {
    //    t.Run(fmt.Sprintf("value for key '%s'", tc.key), func(t *testing.T) {
    usersMap := []map[string]string{
        map[string]string{
            "id": "1",
            "name": "Padraig",
            "email": "padraig@ireland.ie",
            "password": "123safe!",
        },
    }
    users := Users{usersMap}
    search := map[string]string{"name": "Padraig"}
    expected := []*User{}
    user := NewUser("1", "Padraig", "padraig@ireland.ie", "123safe!")
    expected = append(expected, user)
    if res, _ := users.Find(search); res != expected {
        t.Errorf("got %s; want %s", res, expected)
    }
    //    })
    //}
}