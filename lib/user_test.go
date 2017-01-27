package usermanager

import "testing"

func TestNewUser(t *testing.T) {
	User := NewUser("1", "Bob", "bob@baggio.com", "998877abc")

	if User.Id() != "1" {
		t.Error("Expected 1, got ", User.Id())
	}

	if User.Name() != "Bob" {
		t.Error("Expected Bob, got ", User.Name())
	}

	if User.Email() != "bob@baggio.com" {
		t.Error("Expected bob@baggio.com, got ", User.Email())
	}

	if User.Password() != "998877abc" {
		t.Error("Expected 998877abc, got ", User.Password())
	}
}
